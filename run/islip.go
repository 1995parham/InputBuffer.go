/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 26-01-2018
 * |
 * | File Name:     run/islip.go
 * +===============================================
 */

package run

import "github.com/AUTProjects/InputBuffer/algorithm"

func init() {
	m["iSLIP"] = parseISLIP
}

func parseISLIP(parameters map[interface{}]interface{}, n int) (algorithm.Algorithm, error) {
	i := parameters["iterations"].(int)

	alg := algorithm.NewISLIP(n, i)

	alg.AcceptArbiter[0] = 3
	alg.AcceptArbiter[1] = 2
	alg.AcceptArbiter[2] = 2
	alg.AcceptArbiter[3] = 0

	alg.GrantArbiter[0] = 1
	alg.GrantArbiter[1] = 3
	alg.GrantArbiter[2] = 2
	alg.GrantArbiter[3] = 0

	return alg, nil
}
