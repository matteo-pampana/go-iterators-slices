package main

import (
	"fmt"
	"iter"
)

func sliceIteratorWithIndex(slice []int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for index, value := range slice {
			if !yield(index, value) {
				break
			}
		}
	}
}

func sliceIterator(slice []int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, value := range slice {
			if !yield(value) {
				break
			}
		}
	}
}

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, value := range sliceIteratorWithIndex(mySlice) {
		fmt.Println("Index:", index, "Value:", value)
	}

	fmt.Println("--------------------")

	for value := range sliceIterator(mySlice) {
		fmt.Println("Value:", value)
	}
}
