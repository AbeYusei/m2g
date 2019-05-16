package m2g_test

import (
	"testing"

	"github.com/AbeYusei/m2g"
	"github.com/AbeYusei/m2g/testutil"
)

func TestSuccessConvert2Path(t *testing.T) {

	cases := []m2g.Path{
		// {"~/dir.mov", "~/dir", "~/dir.mov", "~/dir.gif"},
		{"~/dir.mov", "~/dir", "~/dir.mov", "~/dir.gif"},
		{"~/dir", "~/dir", "~/dir.mov", "~/dir.gif"},
	}

	for _, tc := range cases {
		if path := m2g.Convert2Path(tc.Raw); tc != *path {
			testutil.Errorf(t, tc, *path)
		}
	}
}
