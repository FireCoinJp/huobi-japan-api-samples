package cmds

// ローソク足

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"
	"net/url"

	"github.com/google/subcommands"
)

type KLineCmd struct {
	symbol string
	period string
	size   string
}

func (a *KLineCmd) Name() string {
	return "kline"
}

func (a *KLineCmd) Synopsis() string {
	return "ローソク足"
}

func (a *KLineCmd) Usage() string {
	return "api-test kline \n"
}

func (a *KLineCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "btcjpy", "取引ペア - btceth")
	set.StringVar(&a.period, "period", "1day", "チャートタイプ")
	set.StringVar(&a.size, "size", "20", "サイズ - default=150 max=2000 ")
}

func (a *KLineCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	param := url.Values{}
	param.Add("symbol", a.symbol)
	param.Add("period", a.period)
	param.Add("size", a.size)

	req, _ := http.NewRequest(http.MethodGet, h.Url("/market/history/kline")+"?"+param.Encode(), nil)

	h.Process(req)
	return 0
}
