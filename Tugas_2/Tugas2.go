package main

import "fmt"

func main() {
	var sisi, keliling, luas int
	fmt.Print("Masukkan sisi :")
	fmt.Scan(&sisi)

	keliling = 4 * sisi
	luas = sisi * sisi

	fmt.Println("Kelilingnya adalah", keliling)
	fmt.Println("Luasnya adalah", luas)

}
