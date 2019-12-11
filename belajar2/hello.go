package main

import "fmt"

func main() {
	//fungsi untuk mencetak teks dengan fmt.Println
	//harus menggunakan fmt karena terdapat import.fmt
	//fmt.Println untuk membuat baris baru
	fmt.Println("test dong")

	//MANIPULASI DATA STRING
	//tuk menghitung jumlah karakter
	fmt.Println(len("apa iya?")) 

	//untuk mengambil karakter tertentu dan outputnya adalah desimal
	fmt.Println("apa iya?"[0])

	//untuk mengambil beberapa huruf
	fmt.Println("goalang"[0:3])
	fmt.Println("golang"[2:])
	fmt.Println("golang"[:4])

	var x string = "selamat pagi"
	fmt.Println(x[0:4])

	/*
		fmt.Print -> tidak membuat baris baru
	*/

}
