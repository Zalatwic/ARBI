//Kevin Szmyd

package EigenCalc

import "fmt"

//normalize the data, O(N'2)
//this means to put it in range [(0,0), (1,1)]
func normalCalc(x [][]float32) [][]float32 {
	var min, max float32 = 999, -999
	var out = make([][]float32, len(x), len(x[0]))

	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[0]); j++ {
			if x[i][j] > max {
				max = x[i][j]
			}

			if x[i][j] < min {
				min = x[i][j]
			}
		}
	}

	//make sure the max value has been modified before division
	max = max - min

	for i := 0; i <= len(x); i++ {
		for j := 0; j <= len(x[0]); j++ {
			out[i][j] = (x[i][j] - min) / max
		}
	}

	return out
}

//create a covariance matrix, O(N'2)
//using yAvg as the average across samples, and xAvg as average per sample
func covCalc(x [][]float32) {
	var yAvg = make([]float32, len(x[0]))
	var out = make([][]float32, len(x), len(x[0]))

	for j := 0; j < len(x[0]); j++ {
		var quickSum float32 = 0

		for i := 0; i < len(x); i++ {
			quickSum += x[i][j]
		}

		yAvg = quickSum
	}

	for i := 0; i < len(x); i++ {
		var quickSum float32 = 0

		for j := 0; j < len(x[0]); j++ {
			quickSum += x[i][j]
		}

		//might want to divide by n - 1 but probably not
		for j := 0; j < len(x[0]); j++ {
			out[i][j] = (x[i][j] - quickSum) * (x[i][j] - yAvg[j])
		}
	}
}
