package main

import (
	"fmt"
	"strconv"
)

func hello(name string) {
	fmt.Println("Hello", name)
}

func sum(a, b int) int {
	return a + b
}

func convertAndSum(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)
	if err != nil {
		return
	}

	total = a + i
	return
}

func main() {
	hello("Anderson")
	println("total:", sum(5, 2))
	total, err := convertAndSum(5, "2")
	println("total:", total, err)
}
