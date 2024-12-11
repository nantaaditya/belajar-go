package main

import "fmt"

func main() {
	var month = [...]string{"january", "february", "march", "april", "mei", "june", "july", "august", "september", "october", "november", "december"}

	var slice = month[4:6]
	fmt.Println(slice)

	month[4] = "mey"
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	var days = [...]string{"senin", "selasa", "rabu", "kami", "jumat", "sabtu", "minggu"}
	var sliceDays = days[5:]
	fmt.Println(sliceDays)

	sliceDays[0] = "sabtu baru"
	fmt.Println(sliceDays)

	newDays := append(sliceDays, "data baru")
	fmt.Println(newDays)

	var dynamicSlice = make([]string, 2, 10)
	dynamicSlice[0] = "1"
	dynamicSlice[1] = "2"
	//dynamicSlice[2] = "3" // error, harus menggunakan append
	fmt.Println(dynamicSlice)

	var dynamicSlice2 = append(dynamicSlice, "3")
	fmt.Println(dynamicSlice2)

	var fromSlice = days[0:]
	var toSlice = make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)
}
