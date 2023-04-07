package ws

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/crypto"
	"huobi-japan-api-samples/data/wsRequest"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type (
	Config struct {
		wsUrl     string
		subs      []string
		path      string
		sleep     time.Duration
		isPrivate bool
		isSave    bool
		accessKey string
		secretKey string
	}

	Builder struct {
		config *Config
		dialer *websocket.Dialer
	}

	Client struct {
		config *Config
		conn   *websocket.Conn
		dialer *websocket.Dialer
	}
)

func wsUrl(host string, isPrivate bool) string {
	if isPrivate {
		return "wss://" + host + "/ws/v2"
	} else {
		return "wss://" + host + "/ws"
	}
}

func NewBuilder(cnf *config.Config, isub wsRequest.ISubscribe) *Builder {
	c := &Config{
		wsUrl:     wsUrl(cnf.Host, isub.IsPrivate()),
		subs:      isub.ToBody(),
		path:      isub.GetPath(),
		isPrivate: isub.IsPrivate(),
		sleep:     cnf.Timeout,
		isSave:    cnf.Save,
		accessKey: cnf.AccessKey,
		secretKey: cnf.SecretKey,
	}
	b := &Builder{
		config: c,
		dialer: websocket.DefaultDialer,
	}

	return b
}

func (b *Builder) Build() *Client {
	var client = &Client{
		config: b.config,
		conn:   nil,
		dialer: b.dialer,
	}

	return client
}

func (w *Client) Run(duration time.Duration) *Client {
	err := w.Connect()
	if err != nil {
		w.systemErrorFunc(err)
		return nil
	}

	go w.ReceiveMessage()
	time.Sleep(duration)
	return nil
}

func (w *Client) handleMessage(msg []byte) error {
	if strings.Contains(string(msg), "ping") {
		pong := strings.Replace(string(msg), "ping", "pong", 1)
		fmt.Println("pong", pong)
		return w.conn.WriteMessage(websocket.TextMessage, []byte(pong))
	}
	var m interface{}
	err := json.Unmarshal(msg, &m)
	if err != nil {
		fmt.Println(err)
	}

	if err := printMsg(m); err != nil {
		return err
	}

	if w.config.isSave {
		err = saveWsMsg(m, w.config.path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *Client) handleAuth() error {
	authParams := url.Values{}
	utc := time.Now().UTC().Format("2006-01-02T15:04:05")
	authParams.Set("accessKey", w.config.accessKey)
	authParams.Set("signatureMethod", "HmacSHA256")
	authParams.Set("signatureVersion", "2.1")
	authParams.Set("timestamp", utc)
	host := "api-cloud.bittrade.co.jp"
	path := "/ws/v2"
	s := fmt.Sprintf("GET\n%s\n%s\n%s", host, path, authParams.Encode())
	signature := crypto.Hmac256(s, w.config.secretKey)
	param := wsRequest.Param{
		AuthType:         "api",
		AccessKey:        w.config.accessKey,
		SignatureMethod:  "HmacSHA256",
		SignatureVersion: "2.1",
		Timestamp:        utc,
		Signature:        signature,
	}
	auth := wsRequest.AuthJson{
		Action: "req",
		Ch:     "auth",
		Params: param,
	}

	authBody, _ := json.Marshal(auth)
	return w.WriteMessage(websocket.TextMessage, authBody)
}

func (w *Client) Connect() error {
	conn, _, err := w.dialer.Dial(w.config.wsUrl, nil)
	if err != nil {
		return err
	}
	w.conn = conn

	if w.config.isPrivate {
		err = w.handleAuth()
		if err != nil {
			return err
		}
	}

	time.Sleep(1 * time.Second)

	if len(w.config.subs) > 0 {
		for _, s := range w.config.subs {
			err = w.conn.WriteMessage(websocket.TextMessage, []byte(s))
			if err != nil {
				w.systemErrorFunc(err)
			}
		}
	}
	return err
}

func (w *Client) ReceiveMessage() {
	var err error
	var msg []byte
	var messageType int
	for {
		messageType, msg, err = w.conn.ReadMessage()
		if err != nil {
			w.systemErrorFunc(err)
			return
		}

		switch messageType {
		case websocket.TextMessage:
			err = w.handleMessage(msg)
			if err != nil {
				w.systemErrorFunc(errors.Wrap(err, "[ws] message handler error."))
			}
		case websocket.BinaryMessage:
			msg, err := w.uncompressFunc(msg)
			if err != nil {
				w.systemErrorFunc(errors.Wrap(err, "[ws] uncompress handler error."))
			} else {
				err = w.handleMessage(msg)
				if err != nil {
					w.systemErrorFunc(errors.Wrap(err, "[ws] uncompress message handler error."))
				}
			}
		}
	}
}

func (w *Client) WriteMessage(messageType int, msg []byte) error {
	return w.conn.WriteMessage(messageType, msg)
}

func (w *Client) uncompressFunc(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r)
}

func (w *Client) systemErrorFunc(err error) {
	fmt.Println("[ws] system error", err.Error())
}

func saveWsMsg(structs interface{}, path string) error {

	fn := "./json/ws/" + path + ".json"
	dir := filepath.Dir(fn)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0700)
		if err != nil {
			return err
		}
	}

	file, _ := os.Create(fn)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(structs)
}

func printMsg(v interface{}) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	err := enc.Encode(v)
	if err != nil {
		return err
	}
	return nil
}
