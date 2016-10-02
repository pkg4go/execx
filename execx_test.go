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

func TestSplit(t *testing.T) {
	a := A{t}

	a.Equal(Split("ls -a -l"), []string{"ls", "-a", "-l"})
	a.Equal(Split("ls -a -l "), []string{"ls", "-a", "-l"})
	a.Equal(Split(" ls -a -l"), []string{"ls", "-a", "-l"})
	a.Equal(Split(" ls -a -l "), []string{"ls", "-a", "-l"})
	a.Equal(Split(" ls   -a   -l "), []string{"ls", "-a", "-l"})

	a.Equal(Split("cmd \"a b c\""), []string{"cmd", "a b c"})
	a.Equal(Split("cmd -a \"a b c\""), []string{"cmd", "-a", "a b c"})
	a.Equal(Split(" cmd -a \"a b c\" "), []string{"cmd", "-a", "a b c"})
	a.Equal(Split(" cmd -a a/b \"a  b  c\" "), []string{"cmd", "-a", "a/b", "a  b  c"})
	a.Equal(Split(" cmd  -a  a/b  \"a  b  c\" "), []string{"cmd", "-a", "a/b", "a  b  c"})
}
