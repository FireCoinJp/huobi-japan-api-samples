package wsRequest

import "encoding/json"

type PrivateRequest struct {
	Action    string `json:"action"`
	Ch        string `json:"ch"`
}

func (p *PrivateRequest) ToBody() []string {
	b, _ := json.Marshal(p)
	return []string{string(b)}
}

func (p *PrivateRequest) GetPath() string {
	return p.Ch
}

func (p *PrivateRequest) IsPrivate() bool {
	return true
}
