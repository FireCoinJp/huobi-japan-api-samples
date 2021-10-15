package api

import (
	"encoding/json"
	"fmt"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/crypto"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

type HandlerResponse func(rep *http.Response) error

type IHttp interface {
	GetAPI(path string, args ...[]string) string
	Url(path string) string
	SetHandler(response HandlerResponse)
	Auth(req *http.Request) error
	Do(req *http.Request, handler HandlerResponse) (*http.Response, error)
}

type Client struct {
	config *config.Config
	client *http.Client
}

func New(cnf *config.Config) *Client {
	return &Client{
		config: cnf,
		client: http.DefaultClient,
	}
}

func (c Client) GetAPI(path string, args ...[]string) string {
	if len(args) > 0 {
		return fmt.Sprintf(c.Url(path), args)
	}
	return c.Url(path)
}

func (c Client) Auth(req *http.Request) error {
	authParams := url.Values{}
	if req.Method == http.MethodGet {
		authParams, _ = url.ParseQuery(req.URL.RawQuery)
	}
	authParams.Set("AccessKeyId", c.config.AccessKey)
	authParams.Set("SignatureMethod", "HmacSHA256")
	authParams.Set("SignatureVersion", "2")
	authParams.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05"))
	s := fmt.Sprintf("%s\n%s\n%s\n%s", req.Method, req.URL.Host, req.URL.Path, authParams.Encode())
	signature := crypto.Hmac256(s, c.config.SecretKey)
	authParams.Set("Signature", signature)
	req.URL, _ = url.Parse(fmt.Sprintf("%s://%s%s?%s", req.URL.Scheme, req.URL.Host, req.URL.Path, authParams.Encode()))
	req.Header.Set("Content-Type", "application/json")
	return nil
}

func (c *Client) Url(path string) string {
	return "https://" + c.config.Host + path
}

func (c Client) Do(req *http.Request, handler HandlerResponse) error {
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	return handler(res)
}

func (c Client) Process(req *http.Request) {
	var err error
	if c.config.Save {
		err = c.Do(req, saveMsg)
	} else {
		err = c.Do(req, printMsg)
	}
	if err != nil {
		panic(err)
	}
}

// handle response functions
func printMsg(rep *http.Response) error {
	var iv interface{}
	err := json.NewDecoder(rep.Body).Decode(&iv)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(iv)
}

func saveMsg(rep *http.Response) error {
	var iv interface{}

	err := json.NewDecoder(rep.Body).Decode(&iv)
	if err != nil {
		return err
	}

	fn := "./json/" + rep.Request.URL.Path + ".json"
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
	return enc.Encode(iv)
}
