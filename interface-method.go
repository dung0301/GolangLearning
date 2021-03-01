package main

import (
	"fmt"
	"math"
)

//khai bao struct
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

type square struct {
	a float64
}

type MyString string

func (myStr MyString) reverse() string {
	s := string(myStr)
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//method
func (r rect) area() float64 {
	return r.width * r.height
}

//method
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (s square) area() float64 {
	return math.Abs(s.a)
}

func (s square) perim() float64 {
	return 4 * s.a
}

//method
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//method
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perim:", g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	s := square{a: 4}

	measure(r)
	measure(c)
	measure(s)

	myStr := MyString("OLLEH")
	fmt.Println(myStr.reverse())
}
