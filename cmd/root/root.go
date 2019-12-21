package root

import (
	"fmt"
	"os"
	"strconv"

	"github.com/calebcase/version/lib/version"
	"github.com/inconshreveable/log15"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
)

var (
	// Log is the logger for the CLI.
	Log = log15.New()

	// RepoPath is the path to the repository.
	RepoPath = "."

	// Cmd is the root command for the CLI.
	Cmd = &cobra.Command{
		Use:   "version",
		Short: "autogenerate versions with patch level",
	}
)

func init() {
	lvl := log15.LvlWarn
	lvlStr, lvlProvided := os.LookupEnv("VERSION_LOG_LEVEL")
	if lvlProvided {
		lvlParsed, err := log15.LvlFromString(lvlStr)
		if err == nil {
			lvl = lvlParsed
		}
	}

	var verbosity uint = 0
	verbosityStr, verbosityProvided := os.LookupEnv("VERSION_LOG_VERBOSITY")
	if verbosityProvided {
		verbosityParsed, err := strconv.ParseUint(verbosityStr, 10, 64)
		if err == nil {
			verbosity = uint(verbosityParsed)
		}
	}

	SetLogger(lvl, verbosity, log15.TerminalFormat())

	flags := Cmd.PersistentFlags()
	flags.StringVarP(&RepoPath, "repopath", "r", RepoPath, "base path for the repository")
}

// SetLogger adjusts the logger Log to with the given log level, verbosity, and
// format.
func SetLogger(lvl log15.Lvl, verbosity uint, format log15.Format) {
	sh := log15.StreamHandler(os.Stderr, format)
	fh := log15.LvlFilterHandler(lvl, sh)

	if verbosity >= 1 {
		fh = log15.CallerFileHandler(fh)
	}

	if verbosity >= 2 {
		fh = log15.CallerFuncHandler(fh)
	}

	if verbosity >= 3 {
		fh = log15.CallerStackHandler("%+v", fh)
	}

	Log.SetHandler(fh)
}

// Repo prints the repository version.
func Repo(vr version.Versioner) (err error) {
	r, err := git.PlainOpen(RepoPath)
	if err != nil {
		return
	}

	v, err := version.Repo(r, vr)
	if err != nil {
		return
	}

	fmt.Println(v)

	return nil
}
