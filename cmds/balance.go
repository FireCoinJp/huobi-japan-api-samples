package cmds

// 残高照合

import (
	"context"
	"flag"
	"fmt"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"

	"github.com/google/subcommands"
)

type BalanceCmd struct {
	isSave bool
}

func (a *BalanceCmd) Name() string {
	return "balance"
}

func (a *BalanceCmd) Synopsis() string {
	return "残高照合"
}

func (a *BalanceCmd) Usage() string {
	return "api-test balance -save"
}

func (a *BalanceCmd) SetFlags(set *flag.FlagSet) {
	set.BoolVar(&a.isSave, "save", false, "write to json")
}

func (a *BalanceCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	req, _ := http.NewRequest(http.MethodGet, h.Url(fmt.Sprintf("/v1/account/accounts/%s/balance", config.Cfg.AccountID)), nil)
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	apiDo(req, a.isSave)
	return 0
}
