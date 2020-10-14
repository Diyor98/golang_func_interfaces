package main

import (
	"fmt"
)

func swap(slice []int, index int) {
	temp := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = temp
}

func bubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				swap(slice, j)
			}
		}
	}
}

func main() {
	mySlice := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Scanln(&mySlice[i])
	}
	bubbleSort(mySlice)
	fmt.Println(mySlice)
}
