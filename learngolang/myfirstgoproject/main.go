package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var elapsed time.Duration
var wg sync.WaitGroup

func main() {

	var c Coffee
	num := 5
	wg.Add(num)
	for i := 0; i < num; i++ {
		go c.CoffeMaker(i)
	}
	wg.Wait()

	fmt.Println("Total time it took to create", num, "coffee:", elapsed.String())

}

type CoffeMethods interface {
	CoffeMaker(i int) string
	grindCoffee(grinder string) string
	makeEspresso(espressoMachine, grounds string)
	steamMilk(steamer string) string
	makeLatte(coffee, milk string)
}

type Coffee struct {
	logger          *log.Logger
	grinder         string
	espressoMachine string
	steamer         string
}

func (c *Coffee) CoffeMaker(i int) {
	c.logger = log.New(os.Stdout, fmt.Sprintf("coffee %d: ", i), log.LstdFlags|log.Lshortfile)
	c.logger.Println("Starting coffee making process")
	start := time.Now()
	grounds := c.grindCoffee(c.grinder)
	coffee := c.makeEspresso(c.espressoMachine, grounds)
	milk := c.steamMilk(c.steamer)
	c.makeLatte(coffee, milk)
	elapsed = time.Since(start)

	defer wg.Done()
}

func (c *Coffee) grindCoffee(grinder string) string {
	c.logger.Println("Grinding the coffee")
	time.Sleep(1 * time.Second)
	return grinder
}

func (c *Coffee) makeEspresso(espressoMachine, grounds string) string {
	c.logger.Println("Making Espresso")
	time.Sleep(1 * time.Second)
	return espressoMachine + grounds
}

func (c *Coffee) steamMilk(steamer string) string {
	c.logger.Println("Steaming the milk")
	time.Sleep(1 * time.Second)
	return steamer
}

func (c *Coffee) makeLatte(coffee, milk string) {
	c.logger.Println("Making the latte")
	time.Sleep(1 * time.Second)
}
