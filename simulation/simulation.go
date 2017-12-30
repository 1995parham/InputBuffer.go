/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-12-2017
 * |
 * | File Name:     simulation.go
 * +===============================================
 */

package simulation

import (
	"fmt"

	"github.com/1995parham/InputBuffer/algorithm"
	"github.com/1995parham/InputBuffer/switches"
)

// Simulation provides switch simulation based on
// given algorithm
type Simulation struct {
	sw  *switches.Switch
	alg algorithm.Algorithm
}

// New creates new simulation instance
func New(sw *switches.Switch, alg algorithm.Algorithm) *Simulation {
	return &Simulation{
		sw:  sw,
		alg: alg,
	}
}

// Simulate run simulation by given number
func (s *Simulation) Simulate(i int) {
	for i > 0 {
		fmt.Printf("%20s%d\n", "I=", i)

		fmt.Println(s.sw)
		m := s.alg.Iterate(s.sw)

		fmt.Println(m)

		for i, o := range m {
			if o != -1 && s.sw.Ports[i].VOQ[o] > 0 {
				s.sw.Ports[i].VOQ[o]--
			}
		}

		i--
	}
}
