package cmds

// ティッカー

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"
	"net/url"

	"github.com/google/subcommands"
)

type MergeCmd struct {
	symbol string
	isSave bool
}

func (a *MergeCmd) Name() string {
	return "merge"
}

func (a *MergeCmd) Synopsis() string {
	return "ティッカー"
}

func (a *MergeCmd) Usage() string {
	return "api-test MergeCmd -save"
}

func (a *MergeCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "btcjpy", "取引ペア")
	set.BoolVar(&a.isSave, "save", false, "write to json")
}

func (a *MergeCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	param := url.Values{}
	param.Add("symbol", a.symbol)

	req, _ := http.NewRequest(http.MethodGet, h.Url("/market/detail/merged")+"?"+param.Encode(), nil)

	apiDo(req, a.isSave)
	return 0
}
