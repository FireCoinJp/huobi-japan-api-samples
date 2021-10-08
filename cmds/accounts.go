package cmds

// ユーザアカウント

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"

	"github.com/google/subcommands"
)

type AccountsCmd struct {
	isSave bool
}

func (a *AccountsCmd) Name() string {
	return "accounts"
}

func (a *AccountsCmd) Synopsis() string {
	return "ユーザアカウント"
}

func (a *AccountsCmd) Usage() string {
	return "api-test accounts -save"
}

func (a *AccountsCmd) SetFlags(set *flag.FlagSet) {
	set.BoolVar(&a.isSave, "save", false, "write to json")
}

func (a *AccountsCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/account/accounts"), nil)
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	apiDo(req, a.isSave)
	return 0
}
