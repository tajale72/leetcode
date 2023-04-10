package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		user := "you : -> "
		fmt.Print(user)
		text, _ := reader.ReadString('\n')

		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if len(text) == 1 {
			TeachAlphabets(text)
		}

		if strings.Contains(text, "slice") && strings.Contains(text, "array") && strings.Contains(text, "what") {
			text = "sliceandarray"
			FindtheFilename(text)
		}

		if strings.Contains(text, "slice") || strings.Contains(text, "array") && strings.Contains(text, "what") {
			FindtheFilename(text)
		}

	}

}

func TeachAlphabets(s string) {
	strMap := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "cherry",
		"d": "dolphin",
		"e": "elephant",
		"f": "flower",
		"g": "giraffe",
		"h": "Home",
		"i": "ice cream",
		"j": "ice cream",
		"k": "kite",
		"l": "lion",
		"m": "monkey",
		"o": "orange",
		"p": "parrot",
		"q": "queen",
		"r": "rabbit",
		"s": "school",
		"t": "table",
		"u": "umbrella",
		"v": "van",
		"w": "water",
		"x": "x-ray",
		"y": "yawn",
		"z": "zebra",
	}
	fmt.Println(s, "for ", strMap[s])
}

func FindtheFilename(text string) {
	AImodel := "AI: -> "
	dir := "/Users/romittajale/Desktop/Desktop/Project1/leetcode/romitAI/answers" // replace with your directory path
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	filename := text + ".txt"
	fmt.Println(filename)
	for _, file := range files {
		if !file.IsDir() && file.Name() == "slice.txt" { // replace with your filename
			filepath := filepath.Join(dir, file.Name())
			data, _ := ioutil.ReadFile(filepath)
			fmt.Println(AImodel, string(data))
		}
	}
}
