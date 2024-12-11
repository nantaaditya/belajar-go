package main

import "fmt"

func main() {
	var name [2]string
	name[0] = "johnson"
	name[1] = "sudrajad"
	fmt.Println(name[0])
	fmt.Println(name[1])

	var number = [3]int{1, 2}
	fmt.Println(number)

	fmt.Println(len(number))

	var month = [...]string{
		"january",
		"february",
		"march",
	}
	fmt.Println(month)
}
