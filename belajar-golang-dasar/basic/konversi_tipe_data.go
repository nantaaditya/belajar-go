package main

import "fmt"

func main() {
	var number32 int32 = 12345
	var number64 int64 = int64(number32)
	fmt.Println(number32)
	fmt.Println(number64)

	var name string = "Hello world"
	var h uint8 = name[0]
	var hString string = string(h)
	fmt.Println(name)
	fmt.Println(h)
	fmt.Println(hString)
}
