package main

import (
	"context"
	"flag"
	"huobi-japan-api-samples/cmds"
	"huobi-japan-api-samples/config"
	"os"

	"github.com/google/subcommands"
)

func init() {
	pwd, _ := os.Getwd()
	var err error
	config.Cfg, err = config.Load(pwd + "/config.yaml")
	if err != nil {
		panic(err)
	}
}

func main() {

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	subcommands.Register(&cmds.SymbolsCmd{}, "システム情報関連")
	subcommands.Register(&cmds.CurrencysCmd{}, "システム情報関連")
	subcommands.Register(&cmds.TimestampCmd{}, "システム情報関連")

	subcommands.Register(&cmds.KLineCmd{}, "マーケット関連")
	subcommands.Register(&cmds.MergeCmd{}, "マーケット関連")
	subcommands.Register(&cmds.TickersCmd{}, "マーケット関連")
	subcommands.Register(&cmds.DepthCmd{}, "マーケット関連")
	subcommands.Register(&cmds.TradeCmd{}, "マーケット関連")
	subcommands.Register(&cmds.HistoryTradeCmd{}, "マーケット関連")

	subcommands.Register(&cmds.AccountsCmd{}, "アカウント関連")
	subcommands.Register(&cmds.BalanceCmd{}, "アカウント関連")

	subcommands.Register(&cmds.PlaceCmd{}, "取引関連")
	subcommands.Register(&cmds.OpenOrdersCmd{}, "取引関連")
	subcommands.Register(&cmds.SubmitcancelCmd{}, "取引関連")
	subcommands.Register(&cmds.BatchcancelCmd{}, "取引関連")
	subcommands.Register(&cmds.BatchCancelOpenOrdersCmd{}, "取引関連")
	subcommands.Register(&cmds.OrderCmd{}, "取引関連")
	subcommands.Register(&cmds.MatchresultsCmd{}, "取引関連")
	subcommands.Register(&cmds.GetOrderCmd{}, "取引関連")
	subcommands.Register(&cmds.GetMatchresultsCmd{}, "取引関連")

	subcommands.Register(&cmds.CreateCmd{}, "ウォレット関連")
	subcommands.Register(&cmds.CancelCmd{}, "ウォレット関連")
	subcommands.Register(&cmds.DepositWithdrawCmd{}, "ウォレット関連")

	subcommands.Register(&cmds.RetailPlaceCmd{}, "販売所関連")
	subcommands.Register(&cmds.OrderlistCmd{}, "販売所関連")
	subcommands.Register(&cmds.MaintainTimeCmd{}, "販売所関連")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
