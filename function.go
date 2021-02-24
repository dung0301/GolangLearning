package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}
func dung(a int, b int, c int) int {
	return a*b + a*c
}
func main() {
	a := sum(1, 2)
	fmt.Println("1+2 =", a)
	b := dung(3, 4, 6)
	fmt.Println("3*4+3*6 =", b)
}
