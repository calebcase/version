package tag

import (
	"github.com/spf13/cobra"

	"github.com/calebcase/version/cmd/root"
	"github.com/calebcase/version/lib/version"
)

var (
	// Cmd is the tag command for the CLI.
	Cmd = &cobra.Command{
		Use:   "tag",
		Short: "determine the base version from the tags",
		Long:  "Determine the base version from the tags and compute the patch level from the number of commits since that tag.",
		RunE: func(command *cobra.Command, args []string) (err error) {
			err = root.Repo(&version.Tag{})

			return
		},
	}
)

func init() {
	root.Cmd.AddCommand(Cmd)
}
