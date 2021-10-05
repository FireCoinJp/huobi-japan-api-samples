package cmds

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"
)

type AccountsCmd struct {
	isSave bool
}

func (a *AccountsCmd) Name() string {
	return "accounts"
}

func (a *AccountsCmd) Synopsis() string {
	return "查询用户账户状态"
}

func (a *AccountsCmd) Usage() string {
	return "api-test accounts -save"
}

func (a *AccountsCmd) SetFlags(set *flag.FlagSet) {
	set.BoolVar(&a.isSave, "save", false, "write to json")
	return
}

func (a *AccountsCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/account/accounts"), nil)
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

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

