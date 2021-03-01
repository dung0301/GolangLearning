package main

import "fmt"

type person struct {
	name string
	age  int
}

// ki tu * dung truoc ten struct
func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println("Person 1:", person{name: "Dung", age: 22})
	fmt.Println("Person 2:", &person{name: "Thao", age: 23})
	s := person{name: "Tho", age: 22}
	fmt.Println("Name of Person 3:", s.name)

	sp := &s
	fmt.Println("Age of Person 3:", sp.age)

	sp.age = 12
	fmt.Println("Age after change", sp.age)
}
