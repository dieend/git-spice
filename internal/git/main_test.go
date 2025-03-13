package git_test

import (
	"testing"

	"github.com/dieend/git-spice/internal/mockedit"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		// mockedit <input>:
		"mockedit": mockedit.Main,
	})
}
