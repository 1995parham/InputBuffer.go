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
		if len(p.voq) != 2 {
			t.Fatalf("Switch input-port VOQ length must be equal to 2 not %d\n", len(p.voq))
		}
	}
}

func TestMatching(t *testing.T) {
	sw := New(2)

	sw.Arrive(0, 0)
	sw.Arrive(0, 1)
	sw.Arrive(1, 1)

	sw.Process(map[int]int{
		0: 0,
		1: 1,
	})

	ps := sw.Process(map[int]int{
		0: 1,
		1: 0,
	})

	if len(ps) != 1 {
		t.Fatalf("There must be one packet but there are/is %d packet[s]", len(ps))
	}
	if sw.t != 2 {
		t.Fatalf("Two timeslots passed from starting not %d", sw.t)
	}
	if ps[0].InputPort != 0 && ps[0].OutputPort != 1 && ps[0].Delay != 1 {
		t.Fatalf("0 -> 1 aftter 1 timeslots != %s", ps[0])
	}
}
