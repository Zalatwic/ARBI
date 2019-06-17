//Kevin Szmyd

package main

import "fmt"
import "LinEng"

func main() {
	test := make([][]float32, 8)
	for x := 0; x < 8; x++ {
		test[x] = make([]float32, 8)
		for y := 0; y < 8; y++ {
			test[x][y] = float32(x + y + 1)
		}
	}

	test = LinEng.NormalCalc(test)
	//test = LinEng.CovCalc(test)

	for x := 0; x < 8; x++ {
		fmt.Println(test[x])
	}
}
