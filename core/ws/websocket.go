package ws

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/data/wsRequest"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/tech-botao/logger"
)

type (
	WsBuilder struct {
		config *WsConfig
		dialer *websocket.Dialer
	}

	WsConfig struct {
		WsUrl             string
		Header            http.Header
		Subs              []string
		IsDump            bool
		IsAutoReconnect   bool
		readDeadLineTime  time.Duration
		reconnectCount    int
		reconnectInterval time.Duration
	}

	WsClient struct {
		ctx    context.Context
		config *WsConfig
		conn   *websocket.Conn
		dialer *websocket.Dialer
		cancel context.CancelFunc

		MessageFunc        func(msg []byte) error
		UncompressFunc     func(msg []byte) ([]byte, error)
		SystemErrorFunc    func(err error)
		AfterConnectedFunc func() error
	}
)
type HandlerResponse func(msg []byte) error

// GetConn 返回Conn的
func (w *WsClient) GetConn() *websocket.Conn {
	return w.conn
}

func NewBuilder() *WsBuilder {
	return &WsBuilder{
		config: &WsConfig{
			IsDump:            false,
			IsAutoReconnect:   false,
			reconnectInterval: time.Second,
			reconnectCount:    100,
		},
		dialer: websocket.DefaultDialer,
	}
}

func (b *WsBuilder) ReconnectCount(count int) *WsBuilder {
	b.config.reconnectCount = count
	return b
}

func (b *WsBuilder) ReconnectInterval(interval time.Duration) *WsBuilder {
	b.config.reconnectInterval = interval
	return b
}

func (b *WsBuilder) URL(url string) *WsBuilder {
	b.config.WsUrl = url
	return b
}

func (b *WsBuilder) Header(header http.Header) *WsBuilder {
	b.config.Header = header
	return b
}

func (b *WsBuilder) Subs(subs []string) *WsBuilder {
	b.config.Subs = subs
	return b
}

func (b *WsBuilder) AutoReconnect() *WsBuilder {
	b.config.IsAutoReconnect = true
	return b
}

func (b *WsBuilder) Dump() *WsBuilder {
	b.config.IsDump = true
	return b
}

func (b *WsBuilder) Dialer(d *websocket.Dialer) *WsBuilder {
	b.dialer = d
	return b
}

func (b *WsBuilder) ReadDeadLineTime(t time.Duration) *WsBuilder {
	b.config.readDeadLineTime = t
	return b
}

func (b *WsBuilder) Build(ctx context.Context, cancel context.CancelFunc) *WsClient {

	var client = &WsClient{
		ctx:    ctx,
		config: b.config,
		conn:   nil,
		dialer: b.dialer,
		cancel: cancel,

		MessageFunc:        DefaultMessageFunc,
		UncompressFunc:     DefaultUncompressFunc,
		SystemErrorFunc:    SystemErrorFunc,
		AfterConnectedFunc: DefaultAfterConnected,
	}
	return client
}

func (b *WsBuilder) New(isub wsRequest.ISubscribe, c *config.Config) *WsClient {
	ctx, cancel := context.WithCancel(context.Background())
	w := b.URL(Url(c.Host, isub.GetIsPrivate())). // TODO
							Subs(isub.ToBody()).
							Dialer(websocket.DefaultDialer).
							Build(ctx, cancel)

	allmsg := make([]interface{}, 0)

	w.MessageFunc = func(msg []byte) error {
		if strings.Contains(string(msg), "ping") {
			pong := strings.Replace(string(msg), "ping", "pong", 1)
			fmt.Println("pong", pong)
			return w.WriteMessage(websocket.TextMessage, []byte(pong))
		}

		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(msg), &m)
		allmsg = append(allmsg, m)
		if err != nil {
			fmt.Println(err)
		}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		fmt.Println("[ws] receive message", enc.Encode(m))

		if c.Save {
			SaveWsMsg(allmsg, isub.GetPath())
		}

		return nil
	}
	if isub.GetIsPrivate() {
		w.AfterConnectedFunc = func() error {
			authParams := url.Values{}
			utc := time.Now().UTC().Format("2006-01-02T15:04:05")
			authParams.Set("accessKey", c.AccessKey)
			authParams.Set("signatureMethod", "HmacSHA256")
			authParams.Set("signatureVersion", "2.1")
			authParams.Set("timestamp", utc)
			host := "api-cloud.huobi.co.jp"
			path := "/ws/v2"
			s := fmt.Sprintf("GET\n%s\n%s\n%s", host, path, authParams.Encode())
			signature := hmac256(s, c.SecretKey)
			param := wsRequest.Param{
				AuthType:         "api",
				AccessKey:        c.AccessKey,
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
	}

	err := w.Connect()
	if err != nil {
		fmt.Println(err)
	}
	go w.ReceiveMessage()

	go func() {
		time.Sleep(200 * time.Second)
		cancel()
	}()
	select {
	case <-ctx.Done():
		logger.Info("exit timeout", time.Now().Format("2006-01-02 15:04:05"))
	}

	return nil
}

func Url(path string, isPrivate bool) string {
	if isPrivate {
		return "wss://" + path + "/ws/v2"
	} else {
		return "wss://" + path + "/ws"
	}
}

func hmac256(base string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(base))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func DefaultUncompressFunc(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r)
}

func DumpResponse(resp *http.Response, body bool) {

	if resp == nil {
		return
	}

	d, _ := httputil.DumpResponse(resp, body)
	fmt.Println("[ws] response:", string(d))
}

func SystemErrorFunc(err error) {
	fmt.Println("[ws] system error", err.Error())
}

func DefaultMessageFunc(msg []byte) error {
	fmt.Println("[ws] receive message", string(msg))
	return nil
}

func DefaultAfterConnected() error {
	fmt.Println("[ws] connect success.", time.Now().Format("2006-01-02 15:04:05"))

	return nil
}

func (w *WsClient) dump(resp *http.Response, body bool) {
	if w.config.IsDump {
		DumpResponse(resp, body)
	}
}

func (w *WsClient) Connect() error {
	conn, resp, err := w.dialer.Dial(w.config.WsUrl, w.config.Header)
	defer w.dump(resp, true)
	if err != nil {
		return err
	}
	w.conn = conn

	err = w.AfterConnectedFunc()
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)

	if len(w.config.Subs) > 0 {
		for _, s := range w.config.Subs {
			err = w.conn.WriteMessage(websocket.TextMessage, []byte(s))
			if err != nil {
				w.SystemErrorFunc(err)
			}
		}
	}
	return err
}

func (w *WsClient) Reconnect() error {
	var err error
	if w.conn != nil {
		err = w.conn.Close()
		if err != nil {
			w.SystemErrorFunc(errors.Wrapf(err, "[ws] [%s] close websocket error", w.config.WsUrl))
			w.conn = nil
			//return err
		}
	}

	for retry := 1; retry <= w.config.reconnectCount; retry++ {
		time.Sleep(w.config.reconnectInterval)
		err = w.Connect()
		if err != nil {
			w.SystemErrorFunc(errors.Wrap(err, fmt.Sprintf("[ws] websocket reconnect fail, retry[%d]", retry)))
			continue
		} else {
			fmt.Println("[ws] retry", retry)
			break
		}
	}

	return err
}

func (w *WsClient) Close() {
	err := w.conn.Close()
	if err != nil {
		w.SystemErrorFunc(errors.Wrapf(err, "[ws] [%s] close websocket error", w.config.WsUrl))
	} else {
		fmt.Println("[ws] close websocket success.", nil)
	}
	time.Sleep(time.Second)
	w.cancel()
}

func (w *WsClient) ReceiveMessage() {
	var err error
	var msg []byte
	var messageType int
	defer w.Close()
	for {
		messageType, msg, err = w.conn.ReadMessage()
		if err != nil {
			w.SystemErrorFunc(err)
			err := w.Reconnect()
			if err != nil {
				w.SystemErrorFunc(errors.Wrap(err, "[ws] quit message loop."))
				return
			}
		}

		// 收到信息后， 延长时间
		if w.config.readDeadLineTime > 0 {
			err = w.conn.SetReadDeadline(time.Now().Add(w.config.readDeadLineTime))
			if err != nil {
				fmt.Println("set readDeadLine error", err)
			}
		}

		switch messageType {
		case websocket.TextMessage:
			err = w.MessageFunc(msg)
			if err != nil {
				w.SystemErrorFunc(errors.Wrap(err, "[ws] message handler error."))
			}
		case websocket.BinaryMessage:
			msg, err := w.UncompressFunc(msg)
			if err != nil {
				w.SystemErrorFunc(errors.Wrap(err, "[ws] uncompress handler error."))
			} else {
				err = w.MessageFunc(msg)
				if err != nil {
					w.SystemErrorFunc(errors.Wrap(err, "[ws] uncompress message handler error."))
				}
			}
		case websocket.CloseAbnormalClosure:
			w.SystemErrorFunc(errors.Wrap(fmt.Errorf("%s", string(msg)), "[ws] abnormal close message"))
		case websocket.CloseMessage:
			w.SystemErrorFunc(errors.Wrap(fmt.Errorf("%s", string(msg)), "[ws] close message"))
		case websocket.CloseGoingAway:
			w.SystemErrorFunc(errors.Wrap(fmt.Errorf("%s", string(msg)), "[ws] goaway message"))
		case websocket.PingMessage:
			fmt.Println("[ws] receive ping", string(msg))
		case websocket.PongMessage:
			fmt.Println("[ws] receive pong", string(msg))
			//case -1: // possibly is noFrame
			//default:
			//	logger.Error(fmt.Sprintf("[ws][%s] error websocket messageType = %d", w.config.WsUrl, messageType), msg)
		}
	}
}

func (w *WsClient) WriteMessage(messageType int, msg []byte) error {
	return w.conn.WriteMessage(messageType, msg)
}

func (w *WsClient) WriteControl(messageType int, msg []byte, deadline time.Time) error {
	return w.conn.WriteControl(messageType, msg, deadline)
}

func (w WsClient) Do(msg []byte, handler HandlerResponse) error {
	// return handler(msg)
	return nil
}

var SaveWsMsg = func(allmsg []interface{}, path string) error {

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
	return enc.Encode(allmsg)
}
