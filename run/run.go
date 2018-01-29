/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 26-01-2018
 * |
 * | File Name:     run/run.go
 * +===============================================
 */

package run

import (
	"fmt"
	"io"
	"time"

	"github.com/AUTProjects/InputBuffer.go/algorithm"
	"github.com/AUTProjects/InputBuffer.go/generator"
	"github.com/AUTProjects/InputBuffer.go/simulation"
	"github.com/AUTProjects/InputBuffer.go/switches"
	"gopkg.in/yaml.v2"
)

// Parser parse given parameters from simulation configuration file
// and return algorithm object with requested configuration or error
type Parser func(parameters map[interface{}]interface{}, n int) (algorithm.Algorithm, error)

var m = map[string]Parser{}

type simulationConfiguration struct {
	Name       string
	Ports      int
	Timeslots  int
	Speedup    int
	Parameters map[interface{}]interface{}
	InputLoad  float64 `yaml:"load"`
	BurstSize  int     `yaml:"burst"`
}

// Run parses simulation configuration string written in yaml format and then run simulatation
// with given configuration
func Run(configuration []byte, w io.Writer) error {
	var s simulationConfiguration

	err := yaml.Unmarshal(configuration, &s)
	if err != nil {
		return err
	}

	if f, ok := m[s.Name]; ok {
		alg, err := f(s.Parameters, s.Ports)
		if err != nil {
			return err
		}

		var p int
		if s.InputLoad > 1.0 || s.InputLoad < 0 {
			p = 80
		} else {
			p = int(s.InputLoad * 100)
		}

		var gen generator.Generator
		if s.BurstSize > 0 {
			// Bursty traffic
			gen = generator.NewBursty(p, s.BurstSize)
		} else {
			// Uniform traffic
			gen = generator.NewUniform(p)
		}

		sw := switches.NewWithSpeedup(s.Ports, s.Speedup)

		end := make(chan int, 1)
		go func() {
			tick := time.Tick(10 * time.Millisecond)
			for {
				select {
				case <-end:
					return
				case <-tick:
					fmt.Printf("%d of %d\n", sw.T(), s.Timeslots)
				}
			}
		}()

		sim := simulation.NewWithWriter(sw, alg, gen, w)
		sim.Simulate(s.Timeslots)

		close(end)
		return nil
	}

	return fmt.Errorf("Invalid algorithm name %s", s.Name)
}
