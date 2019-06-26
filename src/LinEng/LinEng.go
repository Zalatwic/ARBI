//Kevin Szmyd

package LinEng

import "fmt"
import "unsafe"

//normalize the data, O(NH)
func NormalCalc(x [][]float32) [][]float32 {
	var avg, sumt, sd float32 = 0, 0, 0
	var out = make([][]float32, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = make([]float32, len(x[i]))
	}

	//get the average value for the data set
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[0]); j++ {
			avg += x[i][j]
		}
	}

	avg = avg / float32(len(x)*len(x[0]))

	//calculate the standard deviation
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[0]); j++ {
			sumt += (x[i][j] - avg) * (x[i][j] - avg)
		}
	}

	sd = float32(1) / FastInvSqrt(sumt/float32((len(x)*len(x[0]))-1))

	//normalize
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[0]); j++ {
			out[i][j] = (x[i][j] - avg) / sd
		}
	}

	return out
}

//scale the input matrix to range [(0,0), (1,1), O(NH)
func OneScaleCalc(x [][]float32) [][]float32 {
	var min, max float32 = 999, -999
	var out = make([][]float32, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = make([]float32, len(x[i]))
	}

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

	fmt.Println(len(out[0]))

	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[0]); j++ {
			out[i][j] = (x[i][j] - min) / max
		}
	}

	return out
}

//create a covariance matrix, O(N'2)
//using yAvg as the average across samples, and xAvg as average per sample
func CovCalc(x [][]float32) [][]float32 {
	var yAvg = make([]float32, len(x[0]))
	var out = make([][]float32, len(x), len(x[0]))

	for i := 0; i < len(x); i++ {
		out[i] = make([]float32, len(x[0]))
	}

	for j := 0; j < len(x[0]); j++ {
		var quickSum float32 = 0

		for i := 0; i < len(x); i++ {
			quickSum += x[i][j]
		}

		yAvg[j] = quickSum / float32(len(x))
	}

	fmt.Println(yAvg)

	for i := 0; i < len(x[0]); i++ {
		for j := i; j < len(x[0]); j++ {
			var xTemp float32 = 0

			if j == i {
				//calculate variance
				for k := 0; k < len(x); k++ {
					xTemp += (x[k][j] - yAvg[j]) * (x[k][j] - yAvg[j])
				}

				out[j][j] = xTemp / float32(len(x)-1)
			} else {
				//calculate covariance
				for k := 0; k < len(x); k++ {
					xTemp += (x[k][i] - yAvg[i]) * (x[k][j] - yAvg[j])
				}

				out[j][i] = xTemp / float32(len(x)-1)
				out[i][j] = xTemp / float32(len(x)-1)
			}
		}
	}

	return out
}

//flips, O(NH)
func Flip(x [][]float32) [][]float32 {
	var out = make([][]float32, len(x[0]))
	for i := 0; i < len(x[0]); i++ {
		out[i] = make([]float32, len(x))
		for j := 0; j < len(x); j++ {
			out[i][j] = x[j][i]
		}
	}

	return out
}

//generate identity matrix, O(N'2)
func IdentGen(x int) [][]float32 {
	var out = make([][]float32, x)
	for i := 0; i < x; i++ {
		out[i] = make([]float32, x)
		for j := 0; i < x; j++ {
			out[i][j] = 0

			if i == j {
				out[i][j] = 1
			}
		}
	}

	return out
}

//multiply matrix diagonal by float, O(N)
func DiagMulti(x float32, y [][]float32) [][]float32 {
	var out = make([][]float32, len(y))
	for i := 0; i < len(y); i++ {
		out[i] = make([]float32, len(y))
		out[i] = y[i]
		out[i][i] = x * y[i][i]
	}

	return out
}

//subtract matrix x from y, O(N'3)
func MatrixMulti(x [][]float32, y [][]float32) [][]float32 {
	var out = make([][]float32, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = make([]float32, len(y[0]))
		for j := 0; j < len(y[0]); j++ {
			//each i, j component of the output matrix corresponds to the sum of x,ik * y,kj
			//(which matches the out matrix size)
			var temp = float32(0)

			for k := 0; k < len(x); k++ {
				temp += x[i][k] * y[k][j]
			}

			out[i][j] = temp
		}
	}

	return out
}

//you already know what this is
func FastInvSqrt(x float32) float32 {
	xhalf := float32(0.5) * x
	i := *(*int32)(unsafe.Pointer(&x))
	i = int32(0x5f3759df) - int32(i>>1)
	x = *(*float32)(unsafe.Pointer(&i))

	x = x * (1.5 - (xhalf * x * x))

	return x
}
