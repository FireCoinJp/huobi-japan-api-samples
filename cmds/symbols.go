package cmds

// 取引ペア情報

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"

	"github.com/google/subcommands"
)

type SymbolsCmd struct {
}

func (a *SymbolsCmd) Name() string {
	return "symbols"
}

func (a *SymbolsCmd) Synopsis() string {
	return "取引ペア情報"
}

func (a *SymbolsCmd) Usage() string {
	return "api-test symbols \n"
}

func (a *SymbolsCmd) SetFlags(set *flag.FlagSet) {
}

func (a *SymbolsCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/common/symbols"), nil)

	h.Process(req)
	return 0
}
