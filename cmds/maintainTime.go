package cmds

// 販売所注文履歴

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"

	"github.com/google/subcommands"
)

type MaintainTimeCmd struct {
	isSave bool
}

func (a *MaintainTimeCmd) Name() string {
	return "maintainTime"
}

func (a *MaintainTimeCmd) Synopsis() string {
	return "MaintainTimeCmd"
}

func (a *MaintainTimeCmd) Usage() string {
	return "api-test MaintainTimeCmd -save"
}

func (a *MaintainTimeCmd) SetFlags(set *flag.FlagSet) {
	set.BoolVar(&a.isSave, "save", false, "write to json")
	return
}

func (a *MaintainTimeCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/retail/maintain/time"), nil)
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
