//Kevin Szmyd

package VecEng

import "fmt"

//find orthagonal vectors

//return the principal component of given data, O(
//please send in as x[entry][val]
func CallPrince(x [][]float32) []float32 {
	var max, current, y float32 = 0, 0.01, 0.1
	var out = make([]float32, len(x))

	for i := 0; i < len(x[0]); i++ {
		for current > max {
			max = current
			current = 0

			for j := 0; j < len(x); j++ {
				current += x[j][i]
			}
		}
	}
}
