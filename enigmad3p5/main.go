package main

import ("fmt"
		"strings"
)

func main()  {
	nama := "       bang dani mau makan"
	alamat := "          tangerang"
	// angka := 1
	cutTheString(nama, alamat)
}

//untuk penamaan parameter bebas dan mengikuti apa yang dicetak terlebih dahulu
//value yang dipotong mengikuti apayg ada di main
//func cutTheString(kucing string, ayam string) -> urutan tampilan pada output

func cutTheString(kucing string, ayam string)  {
	fmt.Println(strings.TrimSpace(ayam))
	fmt.Println(strings.TrimSpace(kucing))
}