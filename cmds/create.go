package cmds

// 暗号資産の出金申請

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"huobi-japan-api-samples/data/request"
	"net/http"

	"github.com/google/subcommands"
)

type CreateCmd struct {
	address  string
	amount   string
	currency string
	fee      string
	addr_tag string
}

func (a *CreateCmd) Name() string {
	return "create"
}

func (a *CreateCmd) Synopsis() string {
	return "暗号資産の出金申請"
}

func (a *CreateCmd) Usage() string {
	return "api-test create \n"
}

func (a *CreateCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.address, "address", "", "出金アドレス")
	set.StringVar(&a.amount, "amount", "0.6", "出金数量")
	set.StringVar(&a.currency, "currency", "xrp", "通貨種別")
	set.StringVar(&a.fee, "fee", "0.1", "送金手数料")
	set.StringVar(&a.addr_tag, "addr_tag", "", "暗号資産アドレスの共有tag，xrp")
}

func (a *CreateCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	sendBody := request.CreateBody{
		Address:  a.address,
		Amount:   a.amount,
		Currency: a.currency,
		Fee:      a.fee,
		AddrTag:  a.addr_tag,
	}

	placeBody, _ := json.Marshal(sendBody)

	req, _ := http.NewRequest(http.MethodPost, h.Url("/v1/dw/withdraw/api/create"), bytes.NewReader(placeBody))
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	h.Process(req)
	return 0
}
