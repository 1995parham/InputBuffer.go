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
	"io"
	"os"

	"github.com/AUTProjects/InputBuffer.go/algorithm"
	"github.com/AUTProjects/InputBuffer.go/generator"
	"github.com/AUTProjects/InputBuffer.go/switches"
)

// Simulation provides switch simulation based on
// given algorithm
type Simulation struct {
	sw  *switches.Switch
	alg algorithm.Algorithm
	gen generator.Generator
	w   io.Writer
}

// New creates new simulation instance
func New(sw *switches.Switch, alg algorithm.Algorithm, gen generator.Generator) *Simulation {
	return &Simulation{
		sw:  sw,
		alg: alg,
		gen: gen,
		w:   os.Stdout,
	}
}

// NewWithWriter creates new simulation instance with custom writer
func NewWithWriter(sw *switches.Switch, alg algorithm.Algorithm, gen generator.Generator, w io.Writer) *Simulation {
	s := New(sw, alg, gen)
	s.w = w

	return s
}

// Simulate run simulation by given number of time slots
func (s *Simulation) Simulate(timeslots int) {
	counter := 0
	throughput := 0.0
	averageDelay := 0.0

	for counter < timeslots {
		fmt.Fprintf(s.w, "T = %d\n", counter)

		// Generate traffic
		in := s.gen.Generate(s.sw)

		// Generate matching
		m := s.alg.Iterate(s.sw)

		// Print switch state
		fmt.Fprintln(s.w, s.sw)
		// Print matching
		fmt.Fprintln(s.w, m)

		// TODO: Adds print information about algorithm

		// Process switch with generate matching
		out := 0
		delay := 0
		for _, p := range s.sw.Process(m) {
			out++
			delay += p.Delay
		}

		counter++

		if in > out {
			throughput += float64(out) / float64(in)
		} else {
			throughput += 1.0
		}

		if out != 0 {
			averageDelay = float64(delay) / float64(out)
		}
		fmt.Fprintf(s.w, "In: %d, Out: %d, Throughput: %g\n", in, out, float64(out)/float64(in))
		fmt.Fprintf(s.w, "Delay: %d\n", delay)
	}

	fmt.Println("--- Results ---")
	fmt.Printf("Average Throughput: %g\n", throughput/float64(timeslots)*100)
	fmt.Printf("Average Delay: %g\n", averageDelay)
}
