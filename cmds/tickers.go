package cmds

// 全取引ペアの相場情報

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"

	"github.com/google/subcommands"
)

type TickersCmd struct {
	isSave bool
}

func (a *TickersCmd) Name() string {
	return "tickers"
}

func (a *TickersCmd) Synopsis() string {
	return "全取引ペアの相場情報"
}

func (a *TickersCmd) Usage() string {
	return "api-test TickersCmd -save"
}

func (a *TickersCmd) SetFlags(set *flag.FlagSet) {
	set.BoolVar(&a.isSave, "save", false, "write to json")
}

func (a *TickersCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	req, _ := http.NewRequest(http.MethodGet, h.Url("/market/tickers"), nil)

	apiDo(req, a.isSave)
	return 0
}
