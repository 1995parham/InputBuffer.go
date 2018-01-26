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

	"github.com/AUTProjects/InputBuffer.go/algorithm"
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
	Parameters map[interface{}]interface{}
}

// Run parses simulation configuration string written in yaml format and then run simulatation
// with given configuration
func Run(configuration []byte) error {
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

		sim := simulation.New(switches.New(s.Ports), alg)
		sim.Simulate(s.Timeslots)
		return nil
	}

	return fmt.Errorf("Invalid algorithm name %s", s.Name)
}
