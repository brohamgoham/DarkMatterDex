package main

import (
	"os"

	"github.com/brohamgoham/darkmatterdex/app"
	"github.com/brohamgoham/darkmatterdex/cmd/darkmatterdexd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
