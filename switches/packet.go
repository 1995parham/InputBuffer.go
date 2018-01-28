/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 28-01-2018
 * |
 * | File Name:     switches/packet.go
 * +===============================================
 */

package switches

import "fmt"

// Packet contains information about network packet that comes into switch
type Packet struct {
	arrived    int
	InputPort  int
	OutputPort int
	Delay      int
}

func (p *Packet) String() string {
	return fmt.Sprintf("%d -> %d affer %d", p.InputPort, p.OutputPort, p.Delay)
}
