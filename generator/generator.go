/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-01-2018
 * |
 * | File Name:     generator/generator.go
 * +===============================================
 */

package generator

import "github.com/1995parham/InputBuffer.go/switches"

// Generator represents traffic generator algorithm
type Generator interface {
	// Generate adds incoming packets into given switch buffers
	// and returns number of generated packets
	Generate(sw *switches.Switch) int
}
