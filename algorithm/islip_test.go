/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 30-12-2017
 * |
 * | File Name:     islip_test.go
 * +===============================================
 */

package algorithm

import (
	"testing"

	"github.com/AUTProjects/InputBuffer.go/switches"
)

//
// Input 1
//   * 1
// * * 2
//     3
//     4
//
// Input 2
//     1
//     2
//     3
//     4
//
// Input 3
//     1
// * * 2
//     3
//   * 4
//
// Input 4
//     1
//     2
//     3
// * * 4
func TestISLIP1(t *testing.T) {
	sw := switches.New(4)

	sw.ArriveMany(1, 0, 0)
	sw.ArriveMany(2, 0, 1)

	sw.ArriveMany(2, 2, 1)
	sw.ArriveMany(1, 2, 3)

	sw.ArriveMany(2, 3, 3)

	// iSLIP-1
	alg := NewISLIP(4, 1)

	alg.AcceptArbiter[0] = 0
	alg.AcceptArbiter[1] = 0
	alg.AcceptArbiter[2] = 0
	alg.AcceptArbiter[3] = 0

	alg.GrantArbiter[0] = 0
	alg.GrantArbiter[1] = 0
	alg.GrantArbiter[2] = 0
	alg.GrantArbiter[3] = 0

	m := alg.Iterate(sw)
	if m[0] != 0 || m[1] != -1 || m[2] != 3 || m[3] != -1 {
		t.Fatalf("Matching must be [0:0 1:-1 2:3 3:-1] but it is %v\n", m)
	}
}

//
// Input 1
//   * 1
// * * 2
//     3
//     4
//
// Input 2
//     1
//   * 2
//     3
//     4
//
// Input 3
//     1
// * * 2
//     3
//   * 4
//
// Input 4
//     1
//     2
//     3
// * * 4
func TestISLIP2(t *testing.T) {
	sw := switches.New(4)

	sw.ArriveMany(1, 0, 0)
	sw.ArriveMany(2, 0, 1)

	sw.ArriveMany(1, 1, 1)

	sw.ArriveMany(2, 2, 1)
	sw.ArriveMany(1, 2, 3)

	sw.ArriveMany(2, 3, 3)

	// iSLIP-2
	alg := NewISLIP(4, 2)

	alg.AcceptArbiter[0] = 0
	alg.AcceptArbiter[1] = 0
	alg.AcceptArbiter[2] = 0
	alg.AcceptArbiter[3] = 0

	alg.GrantArbiter[0] = 0
	alg.GrantArbiter[1] = 0
	alg.GrantArbiter[2] = 0
	alg.GrantArbiter[3] = 0

	m := alg.Iterate(sw)
	if m[0] != 0 || m[1] != 1 || m[2] != 3 || m[3] != -1 {
		t.Fatalf("Matching must be [0:0 1:1 2:3 3:-1] but it is %v\n", m)
	}

}
