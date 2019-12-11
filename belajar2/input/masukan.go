package main

import "fmt"
import "bufio"
import "os"

func main() {
	//input untuk angka

	var angka1, angka2, hasil int
	fmt.Print("Masukkan angka 1 :")
	fmt.Scan(&angka1)

	fmt.Print("Masukkan angka 2 :")
	fmt.Scan(&angka2)

	hasil = angka1 + angka2

	fmt.Println("Angka yang dijumlahkan adalah", hasil)
//input untuk kalimat
	var kal string
	fmt.Print("Maskan kalimat :")
	scanner := bufio.NewScanner(osStdin)
	scanner.Scan()
	kal = scanner.xt()
	fmt.Println("Kalimatyang dimasukan adalah :", kal)
}
