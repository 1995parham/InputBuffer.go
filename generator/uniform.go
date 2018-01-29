/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-01-2018
 * |
 * | File Name:     generator/uniform.go
 * +===============================================
 */

package generator

import (
	"math/rand"

	"github.com/AUTProjects/InputBuffer.go/switches"
)

// Uniform traffic generator
type Uniform struct {
	p int
}

// NewUniform creates uniform traffic generator with
// load p.
func NewUniform(p int) *Uniform {
	return &Uniform{
		p: p,
	}
}

// Generate adds incomming packets into given switch buffers
// with uniform distribution
// and returns number of generated packets
func (u *Uniform) Generate(sw *switches.Switch) int {
	in := 0
	for i := 0; i < sw.N; i++ {
		if rand.Intn(100) < u.p {
			in++
			sw.Arrive(i, rand.Intn(sw.N))
		}
	}
	return in
}
