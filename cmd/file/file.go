package file

import (
	"github.com/spf13/cobra"

	"github.com/calebcase/version/cmd/root"
	"github.com/calebcase/version/lib/version"
)

var (
	// FileName is the file that will be read for the base version and last
	// commit.
	FileName string = "VERSION"

	// Cmd is the file command for the CLI.
	Cmd = &cobra.Command{
		Use:   "file",
		Short: "determine the base version from a file",
		Long:  "Determine the base version from a file and compute the patch level from the number of commits since that file was modified.",
		RunE: func(command *cobra.Command, args []string) (err error) {
			err = root.Repo(&version.File{
				RepoPath: root.RepoPath,
				FileName: FileName,
			})

			return
		},
	}
)

func init() {
	root.Cmd.AddCommand(Cmd)

	flags := Cmd.PersistentFlags()
	flags.StringVarP(&FileName, "filename", "f", FileName, "file name containing the version")
}
