package main

import (
	"fmt"

	a "leetcode/algorithm"
)

func main() {
	arr := []int{3, 6, 7, 9}

	fmt.Println("Sorted A: ", a.BubbleSort(arr))
	fmt.Println("Sorted B: ", a.QuickSort(arr))
	fmt.Println("Sorted C: ", a.InsertionSort(arr))
	fmt.Println("Sorted D: ", a.MergeSort(arr))
	fmt.Println("Sorted E: ", a.Test1Bubblesort1(arr))
	fmt.Println("Sorted F: ", a.MerSortTest(arr))
	var s a.Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println("Valuse os Stack s x: ", s)
	s.Pop()
	s.Pop()

	fmt.Println("Valuse os Stack s x: ", s)

}