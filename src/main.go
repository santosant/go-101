package main

import "fmt"

var (
	name string
	n1   int
	n2   int32
)

func main() {
	mensagem := "Olá, " + name + "!"
	fmt.Println(mensagem)

	var b, f, s = true, 1.2, "Olá"
	fmt.Println(b, f, s)

	var total int
	total++

	fmt.Println(total)

	var x, y = 10, 20
	fmt.Println(x, y)
	y, x = x, y
	fmt.Println(x, y)

}
