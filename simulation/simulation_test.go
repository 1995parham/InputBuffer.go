package simulation

import (
	"github.com/AUTProjects/InputBuffer.go/algorithm"
	"github.com/AUTProjects/InputBuffer.go/switches"
)

func ExampleDRRM() {
	sw := switches.New(4)

	sw.ArriveMany(2, 0, 0)
	sw.ArriveMany(1, 0, 1)

	sw.ArriveMany(1, 1, 0)
	sw.ArriveMany(2, 1, 2)
	sw.ArriveMany(1, 1, 3)

	sw.ArriveMany(1, 2, 1)
	sw.ArriveMany(3, 2, 2)
	sw.ArriveMany(1, 2, 3)

	sw.ArriveMany(1, 3, 1)
	sw.ArriveMany(1, 3, 2)
	sw.ArriveMany(1, 3, 3)

	alg := algorithm.NewDRRM(4)

	alg.RequestArbiter[0] = 0
	alg.RequestArbiter[1] = 1
	alg.RequestArbiter[2] = 2
	alg.RequestArbiter[3] = 0

	alg.GrantArbiter[0] = 1
	alg.GrantArbiter[1] = 3
	alg.GrantArbiter[2] = 2
	alg.GrantArbiter[3] = 0

	s := New(sw, alg, 0)

	s.Simulate(2)
}

func ExampleISLIP() {
	sw := switches.New(4)

	sw.ArriveMany(3, 0, 0)
	sw.ArriveMany(1, 0, 1)
	sw.ArriveMany(2, 0, 2)

	sw.ArriveMany(1, 1, 0)
	sw.ArriveMany(2, 1, 2)
	sw.ArriveMany(1, 1, 3)

	sw.ArriveMany(1, 2, 1)
	sw.ArriveMany(3, 2, 2)
	sw.ArriveMany(1, 2, 3)

	sw.ArriveMany(1, 3, 1)
	sw.ArriveMany(1, 3, 2)
	sw.ArriveMany(1, 3, 3)

	alg := algorithm.NewISLIP(4, 1)

	alg.AcceptArbiter[0] = 3
	alg.AcceptArbiter[1] = 2
	alg.AcceptArbiter[2] = 2
	alg.AcceptArbiter[3] = 0

	alg.GrantArbiter[0] = 1
	alg.GrantArbiter[1] = 3
	alg.GrantArbiter[2] = 2
	alg.GrantArbiter[3] = 0

	s := New(sw, alg, 0)

	s.Simulate(3)

}
