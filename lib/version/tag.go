package version

import (
	"errors"
	"io"
	"path"

	"github.com/blang/semver"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// Tag implements Versioner using a tag in the repository.
type Tag struct {
	v *semver.Version
}

var _ Versioner = (*Tag)(nil)

// LastCommit returns the commit hash containing of the most recent version
// tag. A tag is considered to contain a version if the tag name is parsable as
// a semantic version (https://semver.org/) with an optional leading 'v' (e.g.
// v1.2.3 or 5.3.0).
func (t *Tag) LastCommit(r *git.Repository) (hash plumbing.Hash, err error) {
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

		if t.v == nil || v.GT(*t.v) {
			t.v = v
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

// Version returns the base version found by parsing the tag name.
func (t *Tag) Version() (v *semver.Version, err error) {
	return t.v, nil
}
