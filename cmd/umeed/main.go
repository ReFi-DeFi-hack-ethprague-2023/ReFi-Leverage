package main

import (
	"os"
	"strings"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	umeeapp "github.com/umee-network/umee/v5/app"
	appparams "github.com/umee-network/umee/v5/app/params"
	"github.com/umee-network/umee/v5/cmd/umeed/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, strings.ToUpper(appparams.Name), umeeapp.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
