/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-12-2017
 * |
 * | File Name:     algorithm/algorithm.go
 * +===============================================
 */

package algorithm

import (
	"fmt"

	"github.com/AUTProjects/InputBuffer/switches"
)

// Match represents input to output port matching
// when there is no match for a port use -1.
type Match map[int]int

func (m Match) String() string {
	s := ""

	for i, o := range m {
		if o == -1 {
			s += fmt.Sprintf("%d -> x\n", i)
		} else {
			s += fmt.Sprintf("%d -> %d\n", i, o)
		}
	}

	return s
}

// Algorithm represents switching fabric algorithm
type Algorithm interface {
	Iterate(*switches.Switch) Match // Iterate runs algorithm in order to create match for next time-slot
}
