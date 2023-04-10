package algorithm

func LinearSearch(arr []int, x int) int {
	for i, v := range arr {
		if v == x {
			return i
		}
	}

	return -1
}

func BinarySearch(arr []int, x int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1 // return -1 if element is not found
}

func TestBinarySearch(arr []int, x int) int {
	var left int
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			left = mid + 1
		} else {
			right = mid + 1
		}
	}
	return -1
}
