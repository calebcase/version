package byfile

import (
	"github.com/spf13/cobra"

	"github.com/calebcase/version/cmd/root"
	"github.com/calebcase/version/lib/version"
)

var (
	Cmd = &cobra.Command{
		Use:   "by-tag",
		Short: "determine the base version from the tags",
		Long:  "Determine the base version from the tags and compute the patch level from the number of commits since that tag.",
		RunE: func(command *cobra.Command, args []string) (err error) {
			err = root.Repo(&version.ByTag{})

			return
		},
	}
)

func init() {
	root.Cmd.AddCommand(Cmd)
}
