package cmds

// システム時間を調べる

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"

	"github.com/google/subcommands"
)

type TimestampCmd struct {
}

func (a *TimestampCmd) Name() string {
	return "timestamp"
}

func (a *TimestampCmd) Synopsis() string {
	return "システム時間を調べる"
}

func (a *TimestampCmd) Usage() string {
	return "api-test timestamp \n"
}

func (a *TimestampCmd) SetFlags(set *flag.FlagSet) {
}

func (a *TimestampCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/common/timestamp"), nil)

	h.Process(req)
	return 0
}
