package branch

import (
	"testing"

	"github.com/fbustamant3/go-mutesting/test"
)

func TestMutatorElse(t *testing.T) {
	test.Mutator(
		t,
		MutatorElse,
		"../../testdata/branch/mutateelse.go",
		1,
	)
}
