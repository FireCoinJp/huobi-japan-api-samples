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
	return "板情報"
}

func (a *DepthCmd) Usage() string {
	return "api-test DepthCmd -save"
}

func (a *DepthCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "btcjpy", "取引ペア, 例えば btcjpy")
	set.StringVar(&a.stepType, "type", "step4", "グルーピングレベル, [step0, step1, step2, step3, step4, step5]")
	set.BoolVar(&a.isSave, "save", false, "write to json")
}

func (a *DepthCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	param := url.Values{}
	param.Add("symbol", a.symbol)
	param.Add("type", a.stepType)

	req, _ := http.NewRequest(http.MethodGet, h.Url("/market/depth")+"?"+param.Encode(), nil)

	apiDo(req, a.isSave)
	return 0
}
