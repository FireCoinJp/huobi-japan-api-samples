package wsRequest

import "encoding/json"

type ISubscribe interface {
	ToBody() []string
	GetPath() string
	IsPrivate() bool
}

type PublicRequest struct {
	Sub       string `json:"sub,omitempty"`
	Req       string `json:"req,omitempty"`
	Id        string `json:"id,omitempty"`
}

func (p *PublicRequest) ToBody() []string {
	b, _ := json.Marshal(p)
	return []string{string(b)}
}

func (p *PublicRequest) GetPath() string {
	return p.Sub
}

func (p *PublicRequest) IsPrivate() bool {
	return false
}
