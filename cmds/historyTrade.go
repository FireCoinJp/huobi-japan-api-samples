package cmds

// 取引履歴の取得

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"
	"net/url"

	"github.com/google/subcommands"
)

type HistoryTradeCmd struct {
	symbol string
	size   string
}

func (a *HistoryTradeCmd) Name() string {
	return "historytrade"
}

func (a *HistoryTradeCmd) Synopsis() string {
	return "取引履歴の取得"
}

func (a *HistoryTradeCmd) Usage() string {
	return "api-test historytrade \n"
}

func (a *HistoryTradeCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "btcjpy", "取引ペア, 例えば btcjpy")
	set.StringVar(&a.size, "size", "2", "データサイズ, Range: {1, 2000}")
}

func (a *HistoryTradeCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	param := url.Values{}
	param.Add("symbol", a.symbol)
	param.Add("size", a.size)

	req, _ := http.NewRequest(http.MethodGet, h.Url("/market/history/trade")+"?"+param.Encode(), nil)

	h.Process(req)
	return 0
}
