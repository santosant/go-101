package main

import (
	"fmt"
	"log"
	"os"
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

func conditions(a, b int) {

	if a > b {
		fmt.Println("a is greater than b")
	} else if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is equal to b")
	}

	file, err := os.Open("hello.txt")
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)
	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))

}

func main() {
	hello("World")
	fmt.Println(sum(1, 2))
	fmt.Println(convertAndSum(1, "2"))
	conditions(2, 2)
}
