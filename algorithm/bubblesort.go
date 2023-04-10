package algorithm

//BubbleSort is a sorting algortihm which sorts an arrray
//by comparing the two adjacent element by comapring the two elements
func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func TestBubbleSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func QuickSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	//we need pivot element

	pi := arr[0]
	var left, right []int
	for _, v := range arr[1:] {
		if v <= pi {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)
	return append(append(left, pi), right...)

}

//InsertionSort is a function which
func InsertionSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	for i := 1; i < n; i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}

	return arr

}

func MergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	middle := n / 2

	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]

		} else {
			result = append(result, right[0])
			right = right[1:]
		}

	}

	result = append(result, left...)
	result = append(result, right...)

	return result

}

func TestMergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2

	left := TestMergeSort(arr[:middle])
	right := TestMergeSort(arr[middle:])

	return merge(left, right)
}

func Testmerge(left, right []int) []int {

	result := []int{}

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	result = append(result, left...)
	result = append(result, right...)

	return result
}

func Test1Bubblesort1(arr []int) []int {

	n := len(arr)
	if n <= 1 {
		return arr
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func MerSortTest(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	mid := n / 2
	left := MerSortTest(arr[:mid])
	right := MerSortTest(arr[mid:])

	return mergeTest(left, right)
}

func mergeTest(left, right []int) []int {

	var result []int

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	result = append(result, left...)
	result = append(result, right...)

	return result

}
