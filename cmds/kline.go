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
	isSave bool
}

func (a *KLineCmd) Name() string {
	return "kline"
}

func (a *KLineCmd) Synopsis() string {
	return "KLineCmd"
}

func (a *KLineCmd) Usage() string {
	return "api-test KLineCmd -save"
}

func (a *KLineCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "btcjpy", "symbol success")
	set.StringVar(&a.period, "period", "1day", "period success")
	set.StringVar(&a.size, "size", "20", "size success")
	set.BoolVar(&a.isSave, "save", false, "write to json")
	return
}

func (a *KLineCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	param := url.Values{}
	param.Add("symbol", a.symbol)
	param.Add("period", a.period)
	param.Add("size", a.size)

	req, _ := http.NewRequest(http.MethodGet, h.Url("/market/history/kline")+"?"+param.Encode(), nil)

	var err error
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
