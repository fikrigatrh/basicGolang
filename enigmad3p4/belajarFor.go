package main

import "fmt"

func main() {
	var fruits = [4]string{"apel", "jeruk", "melon", "pisang"}

	//belajar range untuk menghitung jumlah isi array
	//j -> adalah index yang diberi nama j
	for j, fruit := range fruits {
		fmt.Println(j, fruit)
	}

	//penggunaan _ untuk menghilangkan index, dan bisa disebut juga seperti tempat sampah
	// for _, fruit := range fruits {
	// 	fmt.Println(fruit)
	// }

	for i := 0; i < len(fruits); i++ {
		fmt.Println(i,fruits[i])
	} 
}
