package wsRequest

type AuthJson struct {
	Action string `json:"action"`
	Ch     string `json:"ch"`
	Params Param  `json:"params"`
}

type Param struct {
	AuthType         string `json:"authType"`
	AccessKey        string `json:"accessKey"`
	SignatureMethod  string `json:"signatureMethod"`
	SignatureVersion string `json:"signatureVersion"`
	Timestamp        string `json:"timestamp"`
	Signature        string `json:"signature"`
}
