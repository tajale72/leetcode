package main

import (
	"fmt"
	"log"
	"net/http"
)

type Node struct {
	Value int
	next  *Node
}

func main() {
	res, err := http.Get("https://news.ycombinator.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Load the HTML document using goquery
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the top stories on the page and extract their titles
	doc.Find(".titlelink > .storylink").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		fmt.Printf("%d. %s\n", i+1, title)
	})
}
