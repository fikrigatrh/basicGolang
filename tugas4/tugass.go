package main

import "fmt"

func main() {
	var a, b, c = 10, 20, 0

	c = a
	a = b
	b = c

	/* Cara kedua
	a:=10
	b:=20
	b , a = a, b
	*/

	fmt.Println(a)
	fmt.Println(b)

}
