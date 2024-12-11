package main

import "fmt"

func main() {
	var a int = 10
	var b int = 5
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a += 2
	b -= 2
	fmt.Println(a)
	fmt.Println(b)
}
