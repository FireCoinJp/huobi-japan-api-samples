package cmds

// 対応取引通貨

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"

	"github.com/google/subcommands"
)

type CurrencysCmd struct {
	isSave bool
}

func (a *CurrencysCmd) Name() string {
	return "currencys"
}

func (a *CurrencysCmd) Synopsis() string {
	return "対応取引通貨"
}

func (a *CurrencysCmd) Usage() string {
	return "api-test currencys -save"
}

func (a *CurrencysCmd) SetFlags(set *flag.FlagSet) {
	set.BoolVar(&a.isSave, "save", false, "write to json")
}

func (a *CurrencysCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/common/currencys"), nil)

	apiDo(req, a.isSave)
	return 0
}
