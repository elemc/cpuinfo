package cpuinfo_test

import (
	"bytes"
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

func BenchmarkGet(b *testing.B) {
	var prev []byte
	for i := 0; i < b.N; i++ {
		info, err := cpuinfo.Get()
		if err != nil {
			b.Fatal(err)
		}
		sum := info.Sum()
		if !bytes.Equal(sum[:], prev) && prev != nil {
			b.Fatal("Checksums not equal")
		}
		prev = sum[:]
	}
}
