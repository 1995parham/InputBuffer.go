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

import "github.com/1995parham/InputBuffer/switches"

// Match represents input to output port matching
type Match map[int]int

// Algorithm represents switching fabric algorithm
type Algorithm interface {
	Iterate(*switches.Switch) Match
}
