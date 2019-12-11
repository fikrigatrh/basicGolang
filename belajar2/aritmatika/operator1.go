package main

import "fmt"

func main() {
	a := 5
	b := a + 2
	c := a - b
	d := 7 % 3

	//untuk menambah ada 3 cara yaitu:
	// d++
	d = d + 1
	//d += 1

	fmt.Println("c =", c)
	fmt.Println("d =", d)

	//logical operator
	/*
	&& = and
	|| = or
	!= negasi
	*/

	// true && true = t
	//true && false = f
	//false && false = f
	//true || true = t
	//true || false = t
	//false || false = f
	
	condiLeft := true
	condiRight := false
}
