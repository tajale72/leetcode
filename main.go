package main

import "log"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(arr)-2; i++ {
		if i > 0 && arr[i] == arr[i+1] {
			continue
		}

		log.Println(arr[i])
	}
}
