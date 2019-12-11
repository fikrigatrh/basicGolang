package main

import "fmt"

func main() {
	var a int = 5

	//if else
	fmt.Println("Perintah sebelum if")
	if a < 10 && a >= 5 {
		fmt.Println("perintah ini dieksekusi")
	} else {
		fmt.Println("perintah tidak dieksekusi, yang dieksekusi else")
	}
	fmt.Println("perintah setelah if")

	//if else if
	// if a < 10 && a > 5 {
	// 	fmt.Println("perintah ini dieksekusi")
	// } else if a == 5 {
	// 	fmt.Println("perintah tidak dieksekusi, yang dieksekusi adalah else if")
	// } else {
	// 	fmt.Println("final else")
	// }
	// fmt.Println("setelah if")
}
