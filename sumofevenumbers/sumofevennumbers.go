package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	SumOfEvenNumbers(arr)
}

func SumOfEvenNumbers(arr []int) int {
	var sum int
	for _, v := range arr {
		if v%2 == 0 {
			sum += v
		}
	}
	var name Romit
	name = append(name, "romit")
	fmt.Println(sum)
	return sum
}

type Romit []string
