package cmds

// 板情報

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"
	"net/url"

	"github.com/google/subcommands"
)

type DepthCmd struct {
	symbol   string
	stepType string
	isSave   bool
}

func (a *DepthCmd) Name() string {
	return "depth"
}

func (a *DepthCmd) Synopsis() string {
	return "DepthCmd"
}

func (a *DepthCmd) Usage() string {
	return "api-test DepthCmd -save"
}

func (a *DepthCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "btcjpy", "symbol success")
	set.StringVar(&a.stepType, "type", "step4", "type success")
	set.BoolVar(&a.isSave, "save", false, "write to json")
	return
}

func (a *DepthCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	param := url.Values{}
	param.Add("symbol", a.symbol)
	param.Add("type", a.stepType)

	req, _ := http.NewRequest(http.MethodGet, h.Url("/market/depth")+"?"+param.Encode(), nil)

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
