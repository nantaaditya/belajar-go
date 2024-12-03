package main

import "fmt"

func main() {
	var str string
	str = "ini tipe data string"
	fmt.Println(str)

	var str2 = "ini juga"
	fmt.Println(str2)

	str3 := "ini juga sama"
	fmt.Println(str3)

	var (
		firstName = "first"
		lastName  = "last"
	)
	fmt.Println(firstName, lastName)
}
