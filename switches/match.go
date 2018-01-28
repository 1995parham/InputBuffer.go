/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 28-01-2018
 * |
 * | File Name:     switches/match.go
 * +===============================================
 */

package switches

import "fmt"

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
