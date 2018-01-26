/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 25-01-2018
 * |
 * | File Name:     algorithm/remote.go
 * +===============================================
 */

package algorithm

import (
	"github.com/AUTProjects/InputBuffer.go/switches"
	"github.com/powerman/rpc-codec/jsonrpc2"
)

// Remote represents endpoint for remote algorithm
type Remote struct {
	client *jsonrpc2.Client
}

// NewRemote builds new instance of remote alogrithm endpoint
func NewRemote(url string) *Remote {
	c := jsonrpc2.NewHTTPClient(url)
	return &Remote{
		client: c,
	}
}

// Iterate calls remote algorithm on given switch
func (r *Remote) Iterate(sw *switches.Switch) Match {
	m := make(map[int]int)

	r.client.Call("Algorithm.Iterate", *sw, &m)

	return m
}
