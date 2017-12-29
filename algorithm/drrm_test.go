/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 30-12-2017
 * |
 * | File Name:     drrm_test.go
 * +===============================================
 */

package algorithm

import (
	"fmt"
	"testing"

	"github.com/1995parham/InputBuffer/switches"
)

//
// Input 1
// * * 1
//   * 2
//     3
//     4
//
// Input 2
//   * 1
//     2
// * * 3
//   * 4
//
// Input 3
//     1
//     2
// * * 3
//     4
//
// Input 4
//     1
//   * 2
//   * 3
//     4
func TestDRRM(t *testing.T) {
	sw := switches.New(4)

	sw.Ports[0].VOQ[0] = 2
	sw.Ports[0].VOQ[1] = 1

	sw.Ports[1].VOQ[0] = 1
	sw.Ports[1].VOQ[2] = 2
	sw.Ports[1].VOQ[3] = 1

	sw.Ports[2].VOQ[2] = 2

	sw.Ports[3].VOQ[1] = 1
	sw.Ports[3].VOQ[2] = 1

	alg := NewDRRM(4)

	alg.RequestArbiter[0] = 0
	alg.RequestArbiter[1] = 1
	alg.RequestArbiter[2] = 2
	alg.RequestArbiter[3] = 0

	alg.GrantArbiter[0] = 1
	alg.GrantArbiter[1] = 3
	alg.GrantArbiter[2] = 2
	alg.GrantArbiter[3] = 0

	m := alg.Iterate(sw)
	fmt.Println(m)
}
