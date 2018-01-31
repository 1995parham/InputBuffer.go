/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 31-01-2018
 * |
 * | File Name:     switches/json.go
 * +===============================================
 */

package switches

import "encoding/json"

// MarshalJSON marshals switch into valid json
func (sw *Switch) MarshalJSON() ([]byte, error) {
	ports := make([][]int, sw.N)
	for i := 0; i < sw.N; i++ {
		ports[i] = make([]int, sw.N)
		for j := 0; j < sw.N; j++ {
			ports[i][j] = sw.Ports[i].VOQ(j)
		}
	}

	return json.Marshal(struct {
		N     int
		T     int
		Ports [][]int
	}{
		N:     sw.N,
		T:     sw.T(),
		Ports: ports,
	})
}
