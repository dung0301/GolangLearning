package main

import "fmt"

func sumv(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sumv(1, 2)
	sumv(1, 2, 3)
	sumv(1, 2, 3, 4)
	nums := []int{1, 2, 3, 4, 5}
	sumv(nums...)
}
