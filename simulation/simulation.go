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

// Simulate run simulation by given number
func (s *Simulation) Simulate(i int) {
	for i > 0 {
		s.alg.Iterate(s.sw)
		fmt.Println(s.sw)
	}
}
