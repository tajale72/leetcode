package main

import (
	"fmt"
)

func romanToInt(s string) int {
	m := make(map[string]int)
	// Create a slice of key-value pairs
	pairs := []struct {
		key   string
		value int
	}{
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},
	}

	// Add the pairs to the map
	for _, pair := range pairs {
		m[pair.key] = pair.value
	}
	result := 0

	for k, v := range s {
		value, _ := m[string(v)]
		if k < len(s)-1 && value < m[string(s[k+1])] {
			result -= value
		} else {
			result += value
		}
	}

	return result

}

func main() {
	fmt.Println(romanToInt("III"))     // Output: 3
	fmt.Println(romanToInt("IV"))      // Output: 4
	fmt.Println(romanToInt("IX"))      // Output: 9
	fmt.Println(romanToInt("LVIII"))   // Output: 58
	fmt.Println(romanToInt("MCMXCIV")) // Output: 1994
}
