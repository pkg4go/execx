package execx

import . "github.com/pkg4go/assert"
import "strings"
import "testing"

func TestRun(t *testing.T) {
	a := A{t}

	out, err := Run("ls", "-a")
	a.Equal(err, nil)
	a.Equal(strings.Contains(out, "execx.go"), true)
	a.Equal(strings.Contains(out, "execx_test.go"), true)

	out, err = Run("foo")
	a.NotEqual(err, nil)
	a.NotEqual(out, "")
}
