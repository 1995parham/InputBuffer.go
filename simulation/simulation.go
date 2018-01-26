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

	"github.com/AUTProjects/InputBuffer.go/algorithm"
	"github.com/AUTProjects/InputBuffer.go/switches"
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

// Simulate run simulation by given number of time slots
func (s *Simulation) Simulate(timeslots int) {
	counter := 0
	for counter < timeslots {
		fmt.Printf("%20s%d\n", "T=", counter)
		fmt.Println(s.sw)

		// TODO: Let's generate traffic
		fmt.Println("Generate random traffic")

		m := s.alg.Iterate(s.sw)

		fmt.Println(s.sw)
		fmt.Println(m)
		// TODO: Adds print information about algorithm

		for i, o := range m {
			if o != -1 && s.sw.Ports[i].VOQ[o] > 0 {
				s.sw.Ports[i].VOQ[o]--
			}
		}

		timeslots--
	}
}
