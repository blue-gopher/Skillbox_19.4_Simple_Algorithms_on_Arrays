package main

import (
	"fmt"
)

func mergingTwoSortedArrays(arrFourSize [4]int, arrFiveSize [5]int) [9]int {
	var arr [9]int
	var j, k int
	for i := 0; i < 9; i++ {
		if j < 4 && arrFourSize[j] < arrFiveSize[k] {
			arr[i] = arrFourSize[j]
			j++
		} else if k < 5 {
			arr[i] = arrFiveSize[k]
			k++
		}
	}
	return arr
}

func main() {
	arrFourSize := [4]int{1, 3, 6, 8}
	arrFiveSize := [5]int{2, 4, 5, 7, 9}
	resultArr := mergingTwoSortedArrays(arrFourSize, arrFiveSize)
	fmt.Println(resultArr)
}
