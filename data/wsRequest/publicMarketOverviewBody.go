package wsRequest

import "encoding/json"

type PublicMarketOverviewBody struct {
	Req       string `json:"req"`
	Id        string `json:"id"`
	IsPrivate bool   `json:"isPrivate"`
}

func (p *PublicMarketOverviewBody) ToBody() []string {
	b, _ := json.Marshal(p)
	return []string{string(b)}
}

func (p *PublicMarketOverviewBody) GetPath() string {
	return p.Req
}

func (p *PublicMarketOverviewBody) GetIsPrivate() bool {
	return p.IsPrivate
}
