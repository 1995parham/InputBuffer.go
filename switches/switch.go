/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-12-2017
 * |
 * | File Name:     switch.go
 * +===============================================
 */

package switches

// Port represents input port of switch
type Port struct {
	VOQ []int // Virtual Output Queue
}

// Switch represents switching fabric that works
// by input buffer algorithms
type Switch struct {
	Ports []*Port
	N     int
}

// New creates new switch with n in/out ports
func New(n int) *Switch {
	ports := make([]*Port, n)

	for i := 0; i < n; i++ {
		ports[i] = &Port{
			VOQ: make([]int, n),
		}
	}

	return &Switch{
		Ports: ports,
		N:     n,
	}
}
