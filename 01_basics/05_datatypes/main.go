package main

import "fmt"

func main() {
	// Tipedata berbentuk bilangan
	// Bilangan cacah, maksimal hingga 255 : uint
	var three uint = 3
	fmt.Printf("uint datatype example : %d \n", three)

	// Bilangan bulat : int
	var minTwenty int = -20
	fmt.Printf("int datatype example : %d \n", minTwenty)

	// Masih banyak tipedata lainnya yang bisa diterapkan dan tersedia di Go. Bisa dicek disini : https://dasarpemrogramangolang.novalagung.com/A-tipe-data.html

	// Boolean : bool
	var happy bool = true
	fmt.Printf("My feeling now is happy : %t \n", happy)

	// String : string
	var message = `Nama saya "John Doe".
	Salam kenal.
	Mari belajar "Golang".`
	fmt.Println(message)

	/*
		nil, merupakan nilai kosong yang benar-benar kosong.
		Setiap tipedata memiliki cara berbeda untuk mengekspresikan nilai kosong, yang nilai kosong ini bukan nil.
	*/ 
}
