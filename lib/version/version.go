package version

import (
	"errors"
	"io"

	"github.com/blang/semver"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

var ErrLastCommitNotFound = errors.New("version: last commit not found")

type Versioner interface {
	LastCommit(r *git.Repository) (hash plumbing.Hash, err error)
	Version() (v *semver.Version, err error)
}

func CountCommits(r *git.Repository, hash plumbing.Hash) (count uint64, err error) {
	itr, err := r.Log(&git.LogOptions{})
	if err != nil {
		return
	}

	err = itr.ForEach(func(c *object.Commit) (err error) {
		if c.Hash == hash {
			return io.EOF
		}

		count += 1

		return nil
	})
	if err != nil && !errors.Is(err, io.EOF) {
		return
	}

	return count, nil
}

func Repo(r *git.Repository, vr Versioner) (v *semver.Version, err error) {
	lastCommit, err := vr.LastCommit(r)
	if err != nil {
		return
	}

	count, err := CountCommits(r, lastCommit)
	if err != nil {
		return
	}

	v, err = vr.Version()
	if err != nil {
		return
	}

	v.Patch += count

	return v, nil
}
