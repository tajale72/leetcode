package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	
	for i := 1; i < len(strs); i++ {

		for !strings.HasPrefix(strs[i], prefix) {


			prefix = prefix[:len(prefix)-1]
			
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
	fmt.Println(TestLongestPrefix(strs))

}

//TestLongest prefix
func TestLongestPrefix(str []string) string{
	if len(str) <= 1 {
		return ""
	}

	prefix := str[0]

	for i:=1;i<len(str);i++ {
		for !strings.HasPrefix(str[i], prefix) {
			prefix = prefix[:len(prefix)-1]

			if prefix == "" {
				return ""
			}
		}
	}

	return prefix

	
}