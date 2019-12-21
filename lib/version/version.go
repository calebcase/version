package version

import (
	"errors"
	"io"

	"github.com/blang/semver"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// ErrLastCommitNotFound is an error emitted if the last commit cannot be found.
var ErrLastCommitNotFound = errors.New("version: last commit not found")

// Versioner determines the base version and commit hash for the last base
// version update.
type Versioner interface {
	LastCommit(r *git.Repository) (hash plumbing.Hash, err error)
	Version() (v *semver.Version, err error)
}

// CountCommits returns the number of commits in the repo since the specified
// hash.
func CountCommits(r *git.Repository, hash plumbing.Hash) (count uint64, err error) {
	itr, err := r.Log(&git.LogOptions{})
	if err != nil {
		return
	}

	err = itr.ForEach(func(c *object.Commit) (err error) {
		if c.Hash == hash {
			return io.EOF
		}

		count++

		return nil
	})
	if err != nil && !errors.Is(err, io.EOF) {
		return
	}

	return count, nil
}

// Repo returns the version of the repo.
//
// Repo uses the Versioner to determine the base version of the repository and
// adjusts the patch level by the number of commits since the base version was
// last set.
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
