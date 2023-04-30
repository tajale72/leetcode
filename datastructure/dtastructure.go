package main

import "fmt"

func main() {
	Array()
	Slice()
	TestMap()
	TestInterfaces("hello")

	ceo := &Leader{Name: "CEO"}
	coo := &Leader{Name: "COO"}
	cto := &Leader{Name: "CTO"}

	ceo.AddLeader("COO")
	ceo.Next.AddLeader("CTO")

	ceo.AddMember("CFO")
	coo.AddMember("Manager1")
	coo.AddMember("Manager2")
	cto.AddMember("TechLead1")
	cto.AddMember("TechLead2")
	cto.AddMember("TechLead3")

	ceo.PrintLeadershipChain()
}

// An array is used if we have a fix size data type and want to know the location.
// Example : List of student grades
func Array() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
}

// A Slice is used for a dynamic collection of the same data type.
// Example : Shopping cart, collection of tweet, comments in the picture
func Slice() {
	arr := []int{2, 4, 5, 7}
	fmt.Println(arr)
}

// Map is used when you need a key and a value pair for the data
func TestMap() {
	arr := []int{2, 4, 5, 2, 7}
	new := []int{}
	m := make(map[int]int)

	for _, v := range arr {
		m[v]++
	}

	for _, v := range m {
		if v > 1 {
			new = append(new, v)
		}
	}

	fmt.Println("find duplicates", new)

}

type User struct {
	name string
	age  int
}

// Struct are used when you need multiple fields in a single entity.
func TestStruct() {
	user := &User{
		name: "romit",
		age:  25,
	}
	fmt.Println(user)
}

func TestInterfaces(val interface{}) {
	fmt.Println(val)
}

type Member struct {
	Name string
	Next *Member
}

type Leader struct {
	Name   string
	Member *Member
	Next   *Leader
}

func (l *Leader) AddMember(name string) {
	member := &Member{Name: name}
	if l.Member == nil {
		l.Member = member
	} else {
		current := l.Member
		for current.Next != nil {
			current = current.Next
		}
		current.Next = member
	}
}

func (l *Leader) AddLeader(name string) {
	leader := &Leader{Name: name}
	if l.Next == nil {
		l.Next = leader
	} else {
		current := l.Next
		for current.Next != nil {
			current = current.Next
		}
		current.Next = leader
	}
}

func (l *Leader) PrintLeadershipChain(indent int) {
	fmt.Printf("%s\n", l.Name)
	current := l.Member
	for current != nil {
		for i := 0; i < indent; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("%s\n", current.Name)
		current = current.Next
	}
	if l.Next != nil {
		fmt.Println()
		l.Next.PrintLeadershipChain(indent + 2)
	}
}
