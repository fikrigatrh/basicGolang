package main

import "fmt"

func main()
	var i int = 20
	// var f float32 = 15.5
	var s string

	//belajar konversi tipe data yang berbeda untuk dapat dilakukan operasi aritmatika
	//BELAJAR TYPE CASTING
	// fmt.Println(float32(i)+f)
	fmt.Println(i, "ini juga diprint")
	s = fmt.Sprintln(i, "ini juga diprint")

	fmt.Println("isi variabel i adalah", s)



	//penggunaan Printf
	a, b := 100, 200
	fmt.Printf("printf: ini adalah %[2]v dan %[1]v ", a, b)
}