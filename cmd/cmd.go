package cmd

import (
	"os"

	"github.com/inconshreveable/log15"

	_ "github.com/calebcase/version/cmd/file"
	"github.com/calebcase/version/cmd/root"
	_ "github.com/calebcase/version/cmd/tag"
)

func Main() {
	if err := root.Cmd.Execute(); err != nil {
		root.Log.Error("Failed to execute command.", log15.Ctx{
			"err": err,
		})
		os.Exit(1)
	}
}
