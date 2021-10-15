package cmds

// 資産変動

import (
	"context"
	"flag"
	"fmt"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/ws"
	"huobi-japan-api-samples/data/wsRequest"

	"github.com/google/subcommands"
)

type WsAccountsCmd struct {
	mode string
}

func (a *WsAccountsCmd) Name() string {
	return "wsAccounts"
}

func (a *WsAccountsCmd) Synopsis() string {
	return "資産変動"
}

func (a *WsAccountsCmd) Usage() string {
	return "api-test wsAccounts \n"
}

func (a *WsAccountsCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.mode, "mode", "1", "0	残高が変動する時のみ通知される。1	使える残高に変更があったとき、別々のデータを受信されます。2	使える残高に変更があったとき、同じデータを受信されます。")
}

func (a *WsAccountsCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	w := ws.NewBuilder()
	channel := fmt.Sprintf("accounts.update#%s", a.mode)

	sub := &wsRequest.PrivateOrderBody{
		Action:    "sub",
		Ch:        channel,
		IsPrivate: true,
	}
	w.New(sub, config.Cfg)

	return 0
}
