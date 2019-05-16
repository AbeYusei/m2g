package m2g_test

import (
	"testing"

	"github.com/AbeYusei/m2g"
	"github.com/AbeYusei/testutil"
)

func TestSuccessConvert2Path(t *testing.T) {

	cases := []m2g.Path{
		// {"~/dir.mov", "~/dir", "~/dir.mov", "~/dir.gif"},
		{"~/dir.mov", "~/dir", "~/dir.mov", "~/dir.gif"},
		{"~/dir", "~/dir", "~/dir.mov", "~/dir.gif"},
	}

	for tc := range cases {
		path := m2g.Convert2Path(tc.Raw)
		t.Log(path)
		if tc != path {
			testutil.Errorf(t, tc, path)
		}
	}
}
