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

// Packet contains information about network packet that comes into switch
type Packet struct {
	arrived    int
	InputPort  int
	OutputPort int
	Delay      int
}
