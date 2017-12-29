/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-12-2017
 * |
 * | File Name:     switch_test.go
 * +===============================================
 */

package switches

import "testing"

func TestBasic(t *testing.T) {
	sw := New(2)

	if sw.N != 2 {
		t.Fatalf("Switch ports must be equal to 2 not %d\n", sw.N)
	}

	for _, p := range sw.Ports {
		if len(p.VOQ) != 2 {
			t.Fatalf("Switch input-port VOQ length must be equal to 2 not %d\n", len(p.VOQ))
		}
	}
}
