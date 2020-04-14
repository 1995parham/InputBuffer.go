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
	"github.com/1995parham/InputBuffer.go/switches"
)

// Algorithm represents switching fabric algorithm
type Algorithm interface {
	Iterate(*switches.Switch) switches.Match // Iterate runs algorithm in order to create match for next time-slot
}
