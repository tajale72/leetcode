package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	host := []string{"https://google.com", "https://facebook.com", "https://youtube.com", "https://udemy.com"}

	// define the file path
	filePath := "romit.html"

	// open the file for writing
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println("failed to open file:", err)
		return
	}
	defer file.Close()

	// iterate over the inputs and write each one to the file
	for _, v := range host {
		response, err := http.Get(string(v))
		if err != nil {
			log.Println("failed to get API data:", err)
			continue
		}
		defer response.Body.Close()

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println("failed to read response data:", err)
			continue
		}

		_, err = file.Write(data)
		if err != nil {
			log.Println("failed to write data to file:", err)
			continue
		}
	}

}
