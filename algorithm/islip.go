/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 30-12-2017
 * |
 * | File Name:     islip.go
 * +===============================================
 */

package algorithm

import "github.com/AUTProjects/InputBuffer/switches"

// ISLIP represents Iterative Round-Robin with SLIP Matching Algorithm
type ISLIP struct {
	I             int
	AcceptArbiter []int
	GrantArbiter  []int
}

// NewISLIP builds new instance of iSLIP for n port switch that runs i iteration in each time-slot
func NewISLIP(n int, i int) *ISLIP {
	return &ISLIP{
		I:             i,
		AcceptArbiter: make([]int, n),
		GrantArbiter:  make([]int, n),
	}
}

// Iterate runs iSLIP algorithm on given switch
func (s *ISLIP) Iterate(sw *switches.Switch) Match {
	if sw.N != len(s.AcceptArbiter) || sw.N != len(s.GrantArbiter) {
		panic("invalid switch size")
	}

	m := make(map[int]int)
	b := make(map[int]bool) // is output port number i busy?

	// runs `s.I` iteation and update arbiters only in first iteration
	itr := 0

	// Iteration 1
	for itr < s.I {
		// Request phase
		r := make(map[int][]int)
		// For each input port
		for i := 0; i < sw.N; i++ {
			if m[i] == -1 || itr == 0 {
				for j := 0; j < sw.N; j++ {
					if sw.Ports[i].VOQ[j] != 0 {
						r[j] = append(r[j], i)
					}
				}
			}
		}

		// Grant phase
		g := make(map[int][]int)
		// For each output port
		for i := 0; i < sw.N; i++ {
			if !b[i] {
				granted := false
				for j := 0; j < sw.N; j++ {
					for _, in := range r[i] {
						if in == s.GrantArbiter[i] {
							if !granted {
								granted = true
								g[in] = append(g[in], i)
							}
						}
					}
					if granted {
						break
					}
					if itr == 0 {
						s.GrantArbiter[i] = (s.GrantArbiter[i] + 1) % sw.N
					}
				}
			}
		}

		// Accept phase
		// For each input port
		for i := 0; i < sw.N; i++ {
			accepted := false
			for k, j := 0, s.AcceptArbiter[i]; k < sw.N; k, j = k+1, (j+1)%sw.N {
				for _, out := range g[i] {
					if out == j {
						if !accepted {
							accepted = true
							m[i] = out
							b[out] = true
							if itr == 0 {
								s.GrantArbiter[out] = (s.GrantArbiter[out] + 1) % sw.N
								s.AcceptArbiter[i] = (j + 1) % sw.N
							}
						}
					}
				}
			}
			if itr == 0 && !accepted {
				m[i] = -1
			}
		}

		itr++
	}

	return m
}
