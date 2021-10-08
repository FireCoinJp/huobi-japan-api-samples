package cmds

// 最新の取引データ

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"
	"net/url"

	"github.com/google/subcommands"
)

type TradeCmd struct {
	symbol string
	isSave bool
}

func (a *TradeCmd) Name() string {
	return "trade"
}

func (a *TradeCmd) Synopsis() string {
	return "最新の取引データ"
}

func (a *TradeCmd) Usage() string {
	return "api-test TradeCmd -save"
}

func (a *TradeCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "btcjpy", "取引ペア")
	set.BoolVar(&a.isSave, "save", false, "write to json")
}

func (a *TradeCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	param := url.Values{}
	param.Add("symbol", a.symbol)

	req, _ := http.NewRequest(http.MethodGet, h.Url("/market/trade")+"?"+param.Encode(), nil)

	apiDo(req, a.isSave)
	return 0
}
