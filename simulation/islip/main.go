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
	"github.com/AUTProjects/InputBuffer/algorithm"
	"github.com/AUTProjects/InputBuffer/simulation"
	"github.com/AUTProjects/InputBuffer/switches"
)

func main() {
	sw := switches.New(4)

	sw.Ports[0].VOQ[0] = 3
	sw.Ports[0].VOQ[1] = 1
	sw.Ports[0].VOQ[2] = 2

	sw.Ports[1].VOQ[0] = 1
	sw.Ports[1].VOQ[2] = 2
	sw.Ports[1].VOQ[3] = 1

	sw.Ports[2].VOQ[1] = 1
	sw.Ports[2].VOQ[2] = 3
	sw.Ports[2].VOQ[3] = 1

	sw.Ports[3].VOQ[1] = 1
	sw.Ports[3].VOQ[2] = 1
	sw.Ports[3].VOQ[3] = 1

	alg := algorithm.NewISLIP(4, 1)

	alg.AcceptArbiter[0] = 3
	alg.AcceptArbiter[1] = 2
	alg.AcceptArbiter[2] = 2
	alg.AcceptArbiter[3] = 0

	alg.GrantArbiter[0] = 1
	alg.GrantArbiter[1] = 3
	alg.GrantArbiter[2] = 2
	alg.GrantArbiter[3] = 0

	s := simulation.New(sw, alg)

	s.Simulate(3)
}
