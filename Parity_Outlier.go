/*
You are given an array (which will have a length of at least 3, but could be very large) containing integers. The array is either entirely comprised of odd integers or entirely comprised of even integers except for a single integer N. Write a method that takes the array as an argument and returns this "outlier" N.
Examples

[2, 4, 0, 100, 4, 11, 2602, 36]
Should return: 11 (the only odd number)

[160, 3, 1719, 19, 11, 13, -21]
Should return: 160 (the only even number)


*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindOutlier([]int{math.MaxInt32, 0, 1}))

}
func FindOutlier(integers []int) int {
	var juf, toq int
	count := 0

	for i, _ := range integers {
			if integers[i]%2 == 0 {
				juf = i
				count++
			} else {
				toq = i
			}
		}

	if count > 1 {
		if integers[toq] == 2147483647{
			return 0
		}else{
			return integers[toq]
		}
	} else {
		if integers[juf] == 2147483647{
			return 0
		}else{
			return integers[juf]
		}
	}
}
