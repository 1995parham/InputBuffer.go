/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 29-01-2018
 * |
 * | File Name:     run/remote.go
 * +===============================================
 */

package run

import "github.com/1995parham/InputBuffer.go/algorithm"

func init() {
	m["Remote"] = parseRemote
}

func parseRemote(parameters map[interface{}]interface{}, n int) (algorithm.Algorithm, error) {
	u := parameters["url"].(string)

	alg := algorithm.NewRemote(u)

	return alg, nil
}
