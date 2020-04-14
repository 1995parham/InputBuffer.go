/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 28-01-2018
 * |
 * | File Name:     run/rr_lqf.go
 * +===============================================
 */

package run

import "github.com/1995parham/InputBuffer.go/algorithm"

func init() {
	m["RR/LQF"] = parseRRLQF
}

func parseRRLQF(parameters map[interface{}]interface{}, n int) (algorithm.Algorithm, error) {
	i := parameters["iterations"].(int)

	alg := algorithm.NewRRLQF(n, i)

	return alg, nil
}
