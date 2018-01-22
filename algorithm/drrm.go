/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 30-12-2017
 * |
 * | File Name:     drrm.go
 * +===============================================
 */

package algorithm

import "github.com/AUTProjects/InputBuffer/switches"

// DRRM represents Dual Round Robin Matching Algorithm
type DRRM struct {
	RequestArbiter []int
	GrantArbiter   []int
}

// NewDRRM builds new instance of drrm for n port switch
func NewDRRM(n int) *DRRM {
	return &DRRM{
		RequestArbiter: make([]int, n),
		GrantArbiter:   make([]int, n),
	}
}

// Iterate runs DRRM algorithm on given switch
func (d *DRRM) Iterate(sw *switches.Switch) Match {
	if sw.N != len(d.GrantArbiter) || sw.N != len(d.RequestArbiter) {
		panic("invalid switch size")
	}

	m := make(map[int]int)

	// Request phase
	// For each input port
	for i := 0; i < sw.N; i++ {
		requested := false
		for j := 0; j < sw.N; j++ {
			if sw.Ports[i].VOQ[d.RequestArbiter[i]] != 0 {
				m[i] = d.RequestArbiter[i]
				requested = true
				break
			}
			d.RequestArbiter[i] = (d.RequestArbiter[i] + 1) % sw.N
		}
		if !requested {
			m[i] = -1
		}
	}

	// Grant phase
	// For each output port
	for i := 0; i < sw.N; i++ {
		granted := false
		for k, j := 0, d.GrantArbiter[i]; k < sw.N; k, j = k+1, (j+1)%sw.N {
			if o := m[j]; o == i {
				if !granted {
					granted = true
					d.RequestArbiter[j] = (d.RequestArbiter[j] + 1) % sw.N
					d.GrantArbiter[i] = (j + 1) % sw.N
				} else {
					m[j] = -1
				}
			}
		}
	}

	return m
}
