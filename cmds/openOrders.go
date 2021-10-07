package cmds

// 未約定注文一覧

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"
	"net/url"

	"github.com/google/subcommands"
)

type OpenOrdersCmd struct {
	accountId string
	symbol    string
	side      string
	size      string
	isSave    bool
}

func (a *OpenOrdersCmd) Name() string {
	return "openOrders"
}

func (a *OpenOrdersCmd) Synopsis() string {
	return "OpenOrdersCmd"
}

func (a *OpenOrdersCmd) Usage() string {
	return "api-test OpenOrdersCmd -save"
}

func (a *OpenOrdersCmd) SetFlags(set *flag.FlagSet) {
	set.BoolVar(&a.isSave, "save", false, "write to json")
	set.StringVar(&a.accountId, "account_id", config.Cfg.AccountID, "account_id success")
	set.StringVar(&a.symbol, "symbol", "btcjpy", "symbol success")
	set.StringVar(&a.side, "side", "buy", "side success")
	set.StringVar(&a.size, "size", "2", "size success")
	return
}

func (a *OpenOrdersCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	param := url.Values{}
	param.Add("account-id", a.accountId)
	param.Add("symbol", a.symbol)
	param.Add("side", a.side)
	param.Add("size", a.size)

	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/order/openOrders")+"?"+param.Encode(), nil)
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	if a.isSave {
		err = h.Do(req, api.SaveMsg)
	} else {
		err = h.Do(req, api.PrintMsg)
	}

	if err != nil {
		panic(err)
	}
	return 0
}
