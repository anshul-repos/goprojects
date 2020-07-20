package main

import "fmt"

//created typed struct Person
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	// newPerson constructs a new person stuct with given name
	p := person{name: name}
	p.age = 42
	return &p
}

// this is main

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	// pointer
	sp := &s
	fmt.Println(sp.age)
	// new comment added in to file
	sp.age = 51
	fmt.Println(sp.age)
}
