package main

import "fmt"

func bubbleSort(arr *[6]int) [6]int {
	for i := 0; i < 5; i++ {
		isExchange := false
		for j := 0; j < 5-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isExchange = true
			}
		}

		if isExchange == false {
			break
		}
	}
	return *arr
}

func main() {
	var arr [6]int
	for i := 0; i < 6; i++ {
		fmt.Scan(&arr[i])
	}

	bubbleSort(&arr)
	fmt.Println(arr)
}
