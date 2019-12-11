package main

import "fmt"

func main() {
	var panjang, lebar, tinggi, pl, pt, lt, plt, luasPermukaan, volume int
	fmt.Print("Masukkan panjang :")
	fmt.Scan(&panjang)

	fmt.Print("Masukkan lebar :")
	fmt.Scan(&lebar)

	fmt.Print("Masukkan tinggi :")
	fmt.Scan(&tinggi)

	//luas permukaan balok = 2 * (p*l + p*t + l*t)
	pl = panjang * lebar
	pt = panjang * tinggi
	lt = lebar * tinggi
	plt = pl + pt + lt
	luasPermukaan = 2 * plt
	volume = panjang * lebar * tinggi

	fmt.Println("luas permukaan balok adalah", luasPermukaan)
	fmt.Println("Volumenya adalah", volume)
}
