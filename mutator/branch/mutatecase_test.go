package branch

import (
	"testing"

	"github.com/fbustamant3/go-mutesting/test"
)

func TestMutatorCase(t *testing.T) {
	test.Mutator(
		t,
		MutatorCase,
		"../../testdata/branch/mutatecase.go",
		3,
	)
}
