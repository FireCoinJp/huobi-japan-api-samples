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
	symbol string
	side   string
	size   string
}

func (a *OpenOrdersCmd) Name() string {
	return "openOrders"
}

func (a *OpenOrdersCmd) Synopsis() string {
	return "未約定注文一覧"
}

func (a *OpenOrdersCmd) Usage() string {
	return "api-test openOrders \n"
}

func (a *OpenOrdersCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "btcjpy", "取引通貨ペア")
	set.StringVar(&a.side, "side", "buy", "取引方向, Range: {'buy', 'sell'}")
	set.StringVar(&a.size, "size", "2", "必要な記録数")
}

func (a *OpenOrdersCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	param := url.Values{}
	param.Add("account-id", config.Cfg.AccountID)
	param.Add("symbol", a.symbol)
	param.Add("side", a.side)
	param.Add("size", a.size)

	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/order/openOrders")+"?"+param.Encode(), nil)
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	h.Process(req)
	return 0
}
