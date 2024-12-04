package main

import "fmt"

func main() {
	type NIK string
	var nik NIK = "1234567890"
	fmt.Println(nik)
	fmt.Println(NIK("987654321"))
}
