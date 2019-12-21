package version

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/blang/semver"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// File implements Versioner using a file in the repository.
type File struct {
	RepoPath string
	FileName string
}

var _ Versioner = (*File)(nil)

// LastCommit returns the commit hash containing the last change to FileName.
func (f *File) LastCommit(r *git.Repository) (hash plumbing.Hash, err error) {
	iter, err := r.Log(&git.LogOptions{
		FileName: &f.FileName,
	})
	if err != nil {
		return
	}
	defer iter.Close()

	var found bool
	err = iter.ForEach(func(c *object.Commit) (err error) {
		hash = c.Hash
		found = true

		return io.EOF
	})
	if err != nil && !errors.Is(err, io.EOF) {
		return
	}
	if !found {
		return hash, ErrLastCommitNotFound
	}

	return hash, nil
}

// Version returns the base version found by parsing the contents of FileName.
func (f *File) Version() (v *semver.Version, err error) {
	vf, err := os.Open(filepath.Join(f.RepoPath, f.FileName))
	if err != nil {
		return
	}

	line := bufio.NewScanner(vf)
	line.Scan()

	v, err = semver.New(line.Text())
	if err != nil {
		return
	}

	return
}
