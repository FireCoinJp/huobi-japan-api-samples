package wsRequest

import "encoding/json"

type ISubscribe interface {
	ToBody() []string
	GetPath() string
	GetIsPrivate() bool
}

type PublicMarketBody struct {
	Sub       string `json:"sub"`
	Id        string `json:"id"`
	IsPrivate bool   `json:"isPrivate"`
}

func (p *PublicMarketBody) ToBody() []string {
	b, _ := json.Marshal(p)
	return []string{string(b)}
}

func (p *PublicMarketBody) GetPath() string {
	return p.Sub
}

func (p *PublicMarketBody) GetIsPrivate() bool {
	return p.IsPrivate
}
