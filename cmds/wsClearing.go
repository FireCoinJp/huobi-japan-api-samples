package cmds

// 注文状態更新

import (
	"context"
	"flag"
	"fmt"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/ws"
	"huobi-japan-api-samples/data/wsRequest"

	"github.com/google/subcommands"
)

type WsClearingCmd struct {
	symbol string
	mode   string
}

func (a *WsClearingCmd) Name() string {
	return "wsClearing"
}

func (a *WsClearingCmd) Synopsis() string {
	return "注文状態更新"
}

func (a *WsClearingCmd) Usage() string {
	return "api-test wsClearing \n"
}

func (a *WsClearingCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "trxjpy", "取引ペア（ワイルドカード　*　使用可）")
	set.StringVar(&a.mode, "mode", "1", "プッシュモード（0 - 成約時のみ通知,1 - 成約とキャンセル両方プッシュ；デフォルト値：0）")
}

func (a *WsClearingCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	w := ws.NewBuilder()
	channel := fmt.Sprintf("trade.clearing#%s#%s", a.symbol, a.mode)

	sub := &wsRequest.PrivateOrderBody{
		Action:    "sub",
		Ch:        channel,
		IsPrivate: true,
	}
	w.New(sub, config.Cfg)

	return 0
}
