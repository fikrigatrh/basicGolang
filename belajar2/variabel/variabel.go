package main

import "fmt"

func main() {
	//int8 = -125 s/d 126
	//int16 = -

	var angka int
	angka = 100

	var kal string
	kal = "ini string"

	//untuk bilangan berkoma
	var berkoma float32
	berkoma = 1.5

	//hasil dari boolean ini adalah true, karena yang dicetak oleh sistem adalah secara berurutan, dan yang terkakhir adalah true
	var kondisi bool
	kondisi = false
	kondisi = true

	//contoh misal membuat variabel angka dengan nilai 20, maka hasil outputnya angka menjadi 20 bukan 100
	angka = 20

	//membuat variabel lebih cepat
	var tes1, tes2 = 10, 2
	tes3, tes4 := 30, "oh"


	//membuat variabelengan := sehingga tidak perlu menulis syntax var
	//variabel untuk float
	nilai := berkoma

	//untuk boolean
	keadaan := kondisi

	//cara menggunakan viabel dengan const
	//nilai yang menggunakan syntax const tidak boleh berubah atau dikatakan nilai tetap
	const phi = 3.14

	fmt.Println(angka)
	fmt.Println(kal)
	fmt.Println(berkoma)
	fmt.Println(kondisi)
	fmt.Println(nilai)
	fmt.Println(keadaan)
	fmt.Println(phi)
	fmt.Println(tes1)
	fmt.Println(tes2)
	fmt.Println(tes3)
	fmt.Println(tes4)
}
