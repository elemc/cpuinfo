package cpuinfo_test

import (
	"testing"

	"github.com/elemc/cpuinfo"
)

func TestGet(t *testing.T) {
	info, err := cpuinfo.Get()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", info)
	t.Logf("Sum: %2x", info.Sum())
}
