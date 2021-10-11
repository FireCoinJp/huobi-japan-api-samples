package cmds

// 販売所メンテナンス時間

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"

	"github.com/google/subcommands"
)

type MaintainTimeCmd struct {
}

func (a *MaintainTimeCmd) Name() string {
	return "maintainTime"
}

func (a *MaintainTimeCmd) Synopsis() string {
	return "販売所メンテナンス時間"
}

func (a *MaintainTimeCmd) Usage() string {
	return "api-test maintainTime \n"
}

func (a *MaintainTimeCmd) SetFlags(set *flag.FlagSet) {
}

func (a *MaintainTimeCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/retail/maintain/time"), nil)
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	h.Process(req)
	return 0
}
