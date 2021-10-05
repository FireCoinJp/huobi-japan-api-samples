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

}

func (a AccountsCmd) Name() string {
	return "accounts"
}

func (a AccountsCmd) Synopsis() string {
	return "查询用户余额"
}

func (a AccountsCmd) Usage() string {
	return "api-test accounts"
}

func (a AccountsCmd) SetFlags(set *flag.FlagSet) {
	return
}

func (a AccountsCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/account/accounts"), nil)
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	err = h.Do(req, api.PrintMsg)
	if err != nil {
		panic(err)
	}
	return 0
}

