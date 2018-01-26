/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 30-12-2017
 * |
 * | File Name:     main.go
 * +===============================================
 */

package main

import (
	"github.com/AUTProjects/InputBuffer.go/algorithm"
	"github.com/AUTProjects/InputBuffer.go/simulation"
	"github.com/AUTProjects/InputBuffer.go/switches"
)

func main() {
	sw := switches.New(4)

	sw.Ports[0].VOQ[0] = 2
	sw.Ports[0].VOQ[1] = 1

	sw.Ports[1].VOQ[0] = 1
	sw.Ports[1].VOQ[2] = 2
	sw.Ports[1].VOQ[3] = 1

	sw.Ports[2].VOQ[1] = 1
	sw.Ports[2].VOQ[2] = 3
	sw.Ports[2].VOQ[3] = 1

	sw.Ports[3].VOQ[1] = 1
	sw.Ports[3].VOQ[2] = 1
	sw.Ports[3].VOQ[3] = 1

	alg := algorithm.NewDRRM(4)

	alg.RequestArbiter[0] = 0
	alg.RequestArbiter[1] = 1
	alg.RequestArbiter[2] = 2
	alg.RequestArbiter[3] = 0

	alg.GrantArbiter[0] = 1
	alg.GrantArbiter[1] = 3
	alg.GrantArbiter[2] = 2
	alg.GrantArbiter[3] = 0

	s := simulation.New(sw, alg)

	s.Simulate(2)
}
