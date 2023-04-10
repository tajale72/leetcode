package main

import (
	"fmt"
	//"runtime"
	"sync"
	//"sync/atomic"
)

var (
	counter int32          // counter is a variable incremented by all goroutines.
	wg      sync.WaitGroup // wg is used to wait for the program to finish.
)

func main() {
	wg.Add(3) // Add a count of two, one for each goroutine.

	arr := []string{"Python", "Java", "Golang"}

	for _, v := range  arr {
		go increment(v)
		go NotIncrememnt(v)
	}


	wg.Wait() // Wait for the goroutines to finish.
	fmt.Println("Counter:", counter)

}

func increment(name string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.

	for range name {
		
		counter++
	}

}


func NotIncrememnt(name string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.

	for range name {
		
		counter++
	}

}