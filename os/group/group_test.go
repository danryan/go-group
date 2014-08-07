package group

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

func check(t *testing.T) {
	if !implemented {
		t.Skip("user: not implemented; skipping tests")
	}
	switch runtime.GOOS {
	case "linux", "freebsd", "darwin", "windows", "plan9":
		// test supported
	default:
		t.Skipf("user: Lookup not implemented on %q; skipping test", runtime.GOOS)
	}
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

func compareGroup(t *testing.T, want, got *Group) {
	if want.Gid != got.Gid {
		t.Errorf("got Gid=%q; want %q", got.Gid, want.Gid)
	}
	if want.Name != got.Name {
		t.Errorf("got Name=%q; want %q", got.Name, want.Name)
	}
}

func TestLookupGroup(t *testing.T) {
	check(t)

	// Test LookupGroupId on the current user
	want, err := CurrentGroup()
	if err != nil {
		t.Fatalf("CurrentGroup: %v", err)
	}
	got, err := LookupGroupId(want.Gid)
	if err != nil {
		t.Fatalf("LookupGroupId: %v", err)
	}
	compareGroup(t, want, got)

	// Test Lookup by groupname, using the groupname from LookupId
	g, err := LookupGroup(got.Name)
	if err != nil {
		t.Fatalf("Lookup: %v", err)
	}
	compareGroup(t, got, g)
}
