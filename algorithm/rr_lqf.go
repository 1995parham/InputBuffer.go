/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 23-01-2018
 * |
 * | File Name:     rr_lqf.go
 * +===============================================
 */

package algorithm

import (
	"math/rand"
	"time"

	"github.com/1995parham/InputBuffer.go/switches"
)

// RRLQF represents Iterative Round-Robin with Longest Queue First Algorithm
type RRLQF struct {
	I    int
	rand *rand.Rand
}

// NewRRLQF builds new instance of RR/LQF for n port switch that runs i iteration in each time-slot
func NewRRLQF(n int, i int) *RRLQF {
	return &RRLQF{
		I:    i,
		rand: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

// Iterate runs RR/LQF algorithm on given switch
func (r *RRLQF) Iterate(sw *switches.Switch) switches.Match {
	m := make(map[int]int)

	// preferred input-output pair
	// input ports are keys
	pi := make(map[int]int)
	// output ports are keys
	po := make(map[int]int)
	// for each input port
	for i := 0; i < sw.N; i++ {
		// preferred output
		j := (i + sw.T()) % sw.N
		pi[i] = j
		po[j] = i

	}

	iter := 0

	for iter < r.I {
		// Grant phase
		g := make(map[int][]int)
		// for each output port
		for i := 0; i < sw.N; i++ {
			// check output match status
			if iter > 0 {
				matched := false
				for _, j := range m {
					if i == j {
						matched = true
					}
				}
				if matched {
					continue
				}
			}
			// output port i implicitly identifies its preferred input, say j
			j := po[i]
			if sw.Ports[j].VOQ(i) > 0 {
				// output i grants its preferred input j if C(j, i) > 0
				g[j] = append(g[j], i)
			} else {
				// otherwise output i grants the VOQ that has the largest C(j, i), where j = 0, 1, ..., N - 1
				max := 0
				index := 0
				for j := 0; j < sw.N; j++ {
					if sw.Ports[j].VOQ(i) > max {
						max = sw.Ports[j].VOQ(i)
						index = j
					}
					if sw.Ports[j].VOQ(i) == max {
						if r.rand.Intn(2) == 1 {
							index = j
						}
					}
				}
				g[index] = append(g[index], i)
			}
		}

		// Accept phase
		// for each input port
		for i := 0; i < sw.N; i++ {
			// check inpput match status
			if iter == 0 || m[i] == -1 {
				if len(g[i]) == 0 {
					m[i] = -1
				} else {
					m[i] = g[i][0]
					for _, j := range g[i] {
						if j == pi[i] {
							m[i] = j
							break
						}
						if sw.Ports[i].VOQ(j) > sw.Ports[i].VOQ(m[i]) {
							m[i] = j
						}
						if sw.Ports[i].VOQ(j) == sw.Ports[i].VOQ(m[i]) {
							if r.rand.Intn(2) == 1 {
								m[i] = j
							}

						}
					}
				}
			}
		}
		iter++
	}

	return m
}
