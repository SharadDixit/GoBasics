package main

import (
	"fmt"
)
func main() {
	var favNums [5] float64
	favNums[0] = 111
	favNums[1] = 222
	favNums[2] = 333
	favNums[3] = 444
	favNums[4] = 555
	fmt.Println(favNums[3])

	favNums1 := []float64{1, 2, 3, 4, 5}
	for i , value := range favNums1{fmt.Println(value,i)}
	for _ , value := range favNums1{fmt.Println(value)}    // " _ " used in order to overcome not using i that is if index is not needed
}
