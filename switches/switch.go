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

import "fmt"

// Port represents input port of switch
type Port struct {
	VOQ []int // Virtual Output Queue
}

// Switch represents switching fabric that can work
// with each algorithm
type Switch struct {
	Ports []*Port
	N     int
}

func (sw *Switch) String() string {
	s := ""

	s += fmt.Sprintf("==============\n")
	for i := 0; i < sw.N; i++ {
		s += fmt.Sprintf("Input %d:\n", i)
		for j := 0; j < sw.N; j++ {
			s += fmt.Sprintf("%*s\n", sw.Ports[i].VOQ[j], "---")
			for k := 0; k < sw.Ports[i].VOQ[j]; k++ {
				s += fmt.Sprintf("*")
			}
			s += fmt.Sprintf("|\n")
			s += fmt.Sprintf("%*s\n", sw.Ports[i].VOQ[j], "---")
		}
	}
	s += fmt.Sprintf("==============\n")

	return s
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
