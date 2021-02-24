package main

import "fmt"

func vals() (int, float32) {
	return 5, 5.6
}
func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	_, d := vals()
	fmt.Println(d)
}
