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
	"math/rand"

	"github.com/AUTProjects/InputBuffer.go/algorithm"
	"github.com/AUTProjects/InputBuffer.go/switches"
)

// Simulation provides switch simulation based on
// given algorithm
type Simulation struct {
	sw  *switches.Switch
	alg algorithm.Algorithm
	p   int
}

// New creates new simulation instance
func New(sw *switches.Switch, alg algorithm.Algorithm, p int) *Simulation {
	return &Simulation{
		sw:  sw,
		alg: alg,
		p:   p,
	}
}

// Simulate run simulation by given number of time slots
func (s *Simulation) Simulate(timeslots int) {
	counter := 0
	throughput := 0.0

	for counter < timeslots {
		fmt.Printf("%20s%d\n", "T=", counter)
		fmt.Println(s.sw)

		// TODO: Let's generate traffic
		fmt.Println("Generate uniform traffic")
		in := 0
		for i := 0; i < s.sw.N; i++ {
			if rand.Intn(100) < s.p {
				in++
				s.sw.Ports[i].VOQ[rand.Intn(s.sw.N)]++
			}
		}

		m := s.alg.Iterate(s.sw)

		fmt.Println(s.sw)
		fmt.Println(m)
		// TODO: Adds print information about algorithm

		out := 0
		for i, o := range m {
			if o != -1 && s.sw.Ports[i].VOQ[o] > 0 {
				s.sw.Ports[i].VOQ[o]--
				out++
			}
		}

		counter++
		throughput += float64(out) / float64(in)
		fmt.Printf("In: %d, Out: %d, Throughput: %g\n", in, out, float64(out)/float64(in))
	}

	fmt.Printf("Average Throughput: %g\n", throughput/float64(timeslots)*100)
}
