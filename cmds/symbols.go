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
	isSave bool
}

func (a *SymbolsCmd) Name() string {
	return "symbols"
}

func (a *SymbolsCmd) Synopsis() string {
	return "取引ペア情報"
}

func (a *SymbolsCmd) Usage() string {
	return "api-test symbols -save"
}

func (a *SymbolsCmd) SetFlags(set *flag.FlagSet) {
	set.BoolVar(&a.isSave, "save", false, "write to json")
}

func (a *SymbolsCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/common/symbols"), nil)

	apiDo(req, a.isSave)
	return 0
}
