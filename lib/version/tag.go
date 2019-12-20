package version

import (
	"errors"
	"io"
	"path"

	"github.com/blang/semver"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type ByTag struct {
	v *semver.Version
}

var _ Versioner = (*ByTag)(nil)

func (bt *ByTag) LastCommit(r *git.Repository) (hash plumbing.Hash, err error) {
	iter, err := r.Tags()
	if err != nil {
		return
	}
	defer iter.Close()

	var found bool
	err = iter.ForEach(func(ref *plumbing.Reference) (err error) {
		basename := path.Base(ref.Name().String())

		// Strip leading 'v' if it is present.
		if len(basename) > 0 && basename[0] == 'v' {
			basename = basename[1:]
		}

		v, err := semver.New(basename)
		if err != nil {
			// If it wasn't a parsable version, then continue to
			// the next tag.
			return nil
		}

		if bt.v == nil || v.GT(*bt.v) {
			bt.v = v
			hash = ref.Hash()
			found = true
		}

		return nil
	})
	if err != nil && !errors.Is(err, io.EOF) {
		return
	}
	if !found {
		return hash, ErrLastCommitNotFound
	}

	return hash, nil
}

func (bt *ByTag) Version() (v *semver.Version, err error) {
	return bt.v, nil
}
