package main

import "fmt"

func main() {
	/*
		Slice adalah reference elemen array.
		Cara mendeklarasikannya mirip dengan array, bisa diperoleh pula dari proses manipulasi array maupun slice lain
	*/
	var vegetables = []string{"brocoli", "spinach", "lettuce", "celery"}
	fmt.Println(vegetables[0])

	// Slice hasil manipulasi slice lain
	var newVege = vegetables[0:2]
	fmt.Println(newVege)

	// Potongan kode berikut untuk pembuktian bahwa slice merupakan tipedata reference
	// Kode awal
	var aVege = vegetables[0:3]
	var bVege = vegetables[1:4]

	var aaVege = aVege[1:2]
	var baVege = bVege[0:1]

	fmt.Println(vegetables)
	fmt.Println(aVege)
	fmt.Println(bVege)
	fmt.Println(aaVege)
	fmt.Println(baVege)

	// Nilai "spinach" diubah menjadi "radish"
	baVege[0] = "radish"

	fmt.Println(vegetables)
	fmt.Println(aVege)
	fmt.Println(bVege)
	fmt.Println(aaVege)
	fmt.Println(baVege)
	// Akhir kode

	// cap() dipakai untuk menghitung lebar dan kapasitas maksimum slice
	fmt.Printf("\n")
	fmt.Println("cap() vs len() usage")
	fmt.Printf("Length vegetables : %d \n", len(vegetables))
	fmt.Printf("Cap vegetables : %d \n", cap(vegetables))

	fmt.Printf("Length aVege : %d \n", len(aVege))
	fmt.Printf("Cap aVege : %d \n", cap(aVege))

	fmt.Printf("Length bVege : %d \n", len(bVege))
	fmt.Printf("Cap bVege : %d \n", cap(bVege))

	// append() untuk menambahkan elemen ke slice dengan posisi paling akhir
	fmt.Printf("\n")
	fmt.Println("append() usage")
	var cVege = append(vegetables, "cabbage")
	for _, cvege := range cVege {
		fmt.Println(cvege)
	}
}
