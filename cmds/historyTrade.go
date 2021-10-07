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
	isSave bool
}

func (a *HistoryTradeCmd) Name() string {
	return "historytrade"
}

func (a *HistoryTradeCmd) Synopsis() string {
	return "HistoryTradeCmd"
}

func (a *HistoryTradeCmd) Usage() string {
	return "api-test HistoryTradeCmd -save"
}

func (a *HistoryTradeCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "btcjpy", "symbol success")
	set.StringVar(&a.size, "size", "2", "size success")
	set.BoolVar(&a.isSave, "save", false, "write to json")
	return
}

func (a *HistoryTradeCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	param := url.Values{}
	param.Add("symbol", a.symbol)
	param.Add("size", a.size)

	req, _ := http.NewRequest(http.MethodGet, h.Url("/market/history/trade")+"?"+param.Encode(), nil)

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
