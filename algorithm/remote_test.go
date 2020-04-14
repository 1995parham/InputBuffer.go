/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 25-01-2018
 * |
 * | File Name:     algorithm/remote_test.go
 * +===============================================
 */

package algorithm

import (
	"net/http/httptest"
	"net/rpc"
	"testing"

	"github.com/1995parham/InputBuffer.go/switches"
	"github.com/powerman/rpc-codec/jsonrpc2"
)

type null struct{}

func (null) Iterate(sw switches.Switch, m *switches.Match) error {
	*m = make(map[int]int)

	for i := 0; i < sw.N; i++ {
		(*m)[i] = i
	}

	return nil
}

func TestRemote(t *testing.T) {
	rpc := rpc.NewServer()
	rpc.RegisterName("Algorithm", null{})

	h := jsonrpc2.HTTPHandler(rpc)
	s := httptest.NewServer(h)
	defer s.Close()

	sw := switches.New(8)

	r := NewRemote(s.URL)

	m := r.Iterate(sw)

	if len(m) != sw.N {
		t.Fatalf("Matching size is %d instead of %d", len(m), sw.N)
	}
	for i := range m {
		if m[i] != i {
			t.Fatalf("m[%d] -> %d != %d", i, m[i], i)
		}
	}
}
