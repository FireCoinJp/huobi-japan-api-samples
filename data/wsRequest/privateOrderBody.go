package wsRequest

import "encoding/json"

type PrivateOrderBody struct {
	Action    string `json:"action"`
	Ch        string `json:"ch"`
	IsPrivate bool   `json:"isPrivate"`
}

func (p *PrivateOrderBody) ToBody() []string {
	b, _ := json.Marshal(p)
	return []string{string(b)}
}

func (p *PrivateOrderBody) GetPath() string {
	return p.Ch
}

func (p *PrivateOrderBody) GetIsPrivate() bool {
	return p.IsPrivate
}
