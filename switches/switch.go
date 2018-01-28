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

import (
	"fmt"

	lists "github.com/emirpasic/gods/lists"
	dll "github.com/emirpasic/gods/lists/doublylinkedlist"
)

// Port represents input port of switch
type Port struct {
	voq []lists.List // Virtual Output Queue
}

func newPort(n int) *Port {
	p := &Port{
		voq: make([]lists.List, n),
	}

	for i := 0; i < n; i++ {
		p.voq[i] = dll.New()
	}

	return p
}

// VOQ returns length of n th VOQ
func (p *Port) VOQ(n int) int {
	return p.voq[n].Size()
}

// Switch represents switching fabric that can work
// with each algorithm
type Switch struct {
	Ports   []*Port
	N       int
	Speedup int
	t       int
}

// T returns timestamps of switches
func (sw *Switch) T() int {
	return sw.t
}

func (sw *Switch) String() string {
	s := ""

	s += fmt.Sprintf("==============\n")
	for i := 0; i < sw.N; i++ {
		s += fmt.Sprintf("Input %d:\n", i)
		for j := 0; j < sw.N; j++ {
			s += fmt.Sprintf("%*s\n", sw.Ports[i].VOQ(j), "---")
			for k := 0; k < sw.Ports[i].VOQ(j); k++ {
				s += fmt.Sprintf("*")
			}
			s += fmt.Sprintf("|\n")
			s += fmt.Sprintf("%*s\n", sw.Ports[i].VOQ(j), "---")
		}
	}
	s += fmt.Sprintf("==============\n")

	return s
}

// ArriveMany packets into switch
func (sw *Switch) ArriveMany(n int, i int, o int) {
	for p := 0; p < n; p++ {
		sw.Arrive(i, o)
	}
}

// Arrive new packet into switch
func (sw *Switch) Arrive(i int, o int) {
	if i < 0 || i >= sw.N {
		return
	}
	if o < 0 || o >= sw.N {
		return
	}
	sw.Ports[i].voq[o].Add(&Packet{
		arrived:    sw.t,
		InputPort:  i,
		OutputPort: o,
		Delay:      0,
	})
}

// Process takes switch into next timeslot based on given match
func (sw *Switch) Process(m Match) []*Packet {
	ps := make([]*Packet, 0)

	for i, o := range m {
		if o != -1 {
			for p := 0; p < sw.Speedup; p++ {
				if sw.Ports[i].VOQ(o) > 0 {
					pi, _ := sw.Ports[i].voq[o].Get(0)
					pp := pi.(*Packet)
					pp.Delay = sw.t - pp.arrived
					ps = append(ps, pp)
					sw.Ports[i].voq[o].Remove(0)
				}
			}
		}
	}
	sw.t++

	return ps
}

// New creates new switch with n in/out ports
func New(n int) *Switch {
	ports := make([]*Port, n)

	for i := 0; i < n; i++ {
		ports[i] = newPort(n)
	}

	return &Switch{
		Ports:   ports,
		N:       n,
		Speedup: 1,
	}
}

// NewWithSpeedup creates new switch with n in/out port and speedup s
func NewWithSpeedup(n int, s int) *Switch {
	sw := New(n)
	sw.Speedup = s

	return sw
}
