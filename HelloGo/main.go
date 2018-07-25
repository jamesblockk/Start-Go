package main

import (
	"fmt"
)

func main() {

}

func task1() {
	fmt.Println("Hello GO")

	s := sum(1, 3, 5, 3)
	fmt.Println(s)
}

func sum(nums ...int) int {
	var i = 0
	for _, n := range nums {
		i = i + n
	}

	return i
}
