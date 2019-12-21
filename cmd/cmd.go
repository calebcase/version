package cmd

import (
	"os"

	"github.com/inconshreveable/log15"

	"github.com/calebcase/version/cmd/root"

	// The subcommands are imported, but not directly used here. Their
	// package init methods will attach them to the root command.
	_ "github.com/calebcase/version/cmd/file"
	_ "github.com/calebcase/version/cmd/tag"
)

// Main executes the root command.
func Main() {
	if err := root.Cmd.Execute(); err != nil {
		root.Log.Error("Failed to execute command.", log15.Ctx{
			"err": err,
		})
		os.Exit(1)
	}
}
