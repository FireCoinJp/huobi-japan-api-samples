package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"huobi-japan-api-samples/cmds"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"os"
)

var (
	cnf *config.Config
	h *api.Client

)


func init() {
	pwd, _ := os.Getwd()
	var err error
	config.Cfg, err = config.Load(pwd+"/config.yaml")
	if err != nil {
		panic(err)
	}
}

func main() {

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&cmds.AccountsCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
