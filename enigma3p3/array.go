package main

import "fmt"

func main() {
	//array = kumpulan data yang bertipe sama
	//array cara 1
	var names [3]string
	names[0] = "te"
	names[1] = "du"
	names[2] = "to"

	//array cara 2
	var nama = [3]string {"anto", "budi", "wahyu"}

	fmt.Println(names)
	fmt.Println(nama)

	// [...] ->berguna jika kita tidak mengetahui jumlah array
	//hasil dari array akan mengikuti isinya atau value
	var num = [...] int{1,2,3,4}

	//len -> untuk mengetahui jumlah dari isi array
	fmt.Println(num, len(nama))


}
