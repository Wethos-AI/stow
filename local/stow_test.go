package local_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/Wethos-AI/stow"
	"github.com/Wethos-AI/stow/test"
	"github.com/cheekybits/is"
)

func TestStow(t *testing.T) {
	is := is.New(t)

	dir, err := ioutil.TempDir("testdata", "stow")
	is.NoErr(err)
	defer os.RemoveAll(dir)
	cfg := stow.ConfigMap{"path": dir}

	test.All(t, "local", cfg)
}
