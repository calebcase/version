package cmd

import (
	"os"

	"github.com/inconshreveable/log15"

	_ "github.com/calebcase/version/cmd/byfile"
	_ "github.com/calebcase/version/cmd/bytag"
	"github.com/calebcase/version/cmd/root"
)

func Main() {
	if err := root.Cmd.Execute(); err != nil {
		root.Log.Error("Failed to execute command.", log15.Ctx{
			"err": err,
		})
		os.Exit(1)
	}
}
