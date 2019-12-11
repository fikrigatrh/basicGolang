package main

import "fmt"

func main()  {

	const phi = 3.14
	var jari, keliling, luas float32
	fmt.Print("Masukkan jari-jari :")
	fmt.Scan(&jari)

	keliling = 2 * phi * jari
	luas = phi * jari * jari
	fmt.Println("Kelilingnya adalah", keliling)
	fmt.Println("Luasnya adalah", luas)
}
