/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-12-2017
 * |
 * | File Name:     algorithm_test.go
 * +===============================================
 */

package algorithm

import (
	"testing"

	"github.com/AUTProjects/InputBuffer/switches"
)

type dummy struct{}

func (d *dummy) Iterate(s *switches.Switch) Match {
	return map[int]int{
		0: -1,
		1: 1,
	}
}

func TestBasic(t *testing.T) {
	var alg Algorithm
	alg = &dummy{}

	sw := switches.New(2)

	m := alg.Iterate(sw)

	for i := 0; i < 2; i++ {
		if m[i] >= 2 {
			t.Fatalf("Input %d -> Output %d but %d >= 2\n", i, m[i], m[i])
		}
	}
}
