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
	isSave bool
}

func (a *TimestampCmd) Name() string {
	return "timestamp"
}

func (a *TimestampCmd) Synopsis() string {
	return "TimestampCmd"
}

func (a *TimestampCmd) Usage() string {
	return "api-test TimestampCmd -save"
}

func (a *TimestampCmd) SetFlags(set *flag.FlagSet) {
	set.BoolVar(&a.isSave, "save", false, "write to json")
	return
}

func (a *TimestampCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/common/timestamp"), nil)

	var err error

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
