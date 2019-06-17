//Kevin Szmyd

package LinEng

import "fmt"
import "unsafe"

//normalize the data, O(N'2)
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

//scale the input matrix to range [(0,0), (1,1), O(N'2)
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
			fmt.Println("i", i)
			fmt.Println("j", j)
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
		out[i] = make([]float32, len(x[i]))
	}

	for j := 0; j < len(x[0]); j++ {
		var quickSum float32 = 0

		for i := 0; i < len(x); i++ {
			quickSum += x[i][j]
		}

		yAvg[j] = quickSum
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
