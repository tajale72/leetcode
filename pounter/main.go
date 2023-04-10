package main

import "fmt"

func main() {
	var user Name

	user.UpdateName()

	fmt.Println(&user)
}

type Name struct {
	name string
}

func (p *Name) UpdateName() {
	p.name = "romit"
}
