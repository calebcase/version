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

type ByFile struct {
	RepoPath string
	FileName string
}

var _ Versioner = (*ByFile)(nil)

func (bf *ByFile) LastCommit(r *git.Repository) (hash plumbing.Hash, err error) {
	iter, err := r.Log(&git.LogOptions{
		FileName: &bf.FileName,
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

func (bf *ByFile) Version() (v *semver.Version, err error) {
	f, err := os.Open(filepath.Join(bf.RepoPath, bf.FileName))
	if err != nil {
		return
	}

	line := bufio.NewScanner(f)
	line.Scan()

	v, err = semver.New(line.Text())
	if err != nil {
		return
	}

	return
}
