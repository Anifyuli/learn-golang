package main

import "fmt"

func main() {
	// Slice adalah reference elemen array.
	// Cara mendeklarasikannya mirip dengan array, bisa diperoleh pula dari proses manipulasi array maupun slice lain
	var vegetables = []string{"brokoli", "bayam", "selada", "seledri"}
	fmt.Println(vegetables[0]) // Output: brokoli

	// Slice hasil manipulasi slice lain
	var newVege = vegetables[0:2]
	fmt.Println(newVege) // Output: [brokoli bayam]

	// Potongan kode berikut untuk pembuktian bahwa slice merupakan tipe data reference
	var aVege = vegetables[0:3]
	var bVege = vegetables[1:4]
	var aaVege = aVege[1:2]
	var baVege = bVege[0:1]

	fmt.Println(vegetables) // Output: [brokoli bayam selada seledri]
	fmt.Println(aVege)      // Output: [brokoli bayam selada]
	fmt.Println(bVege)      // Output: [bayam selada seledri]
	fmt.Println(aaVege)     // Output: [bayam]
	fmt.Println(baVege)     // Output: [bayam]

	// Nilai "bayam" diubah menjadi "lobak"
	baVege[0] = "lobak"

	fmt.Println(vegetables) // Output: [brokoli lobak selada seledri]
	fmt.Println(aVege)      // Output: [brokoli lobak selada]
	fmt.Println(bVege)      // Output: [lobak selada seledri]
	fmt.Println(aaVege)     // Output: [lobak]
	fmt.Println(baVege)     // Output: [lobak]

	// Menggunakan cap() untuk menghitung lebar dan kapasitas maksimum slice
	fmt.Println("\ncap() vs len() usage")
	fmt.Printf("Length vegetables : %d \n", len(vegetables)) // Output: 4
	fmt.Printf("Cap vegetables : %d \n", cap(vegetables))     // Output: 4
	fmt.Printf("Length aVege : %d \n", len(aVege))            // Output: 3
	fmt.Printf("Cap aVege : %d \n", cap(aVege))               // Output: 4
	fmt.Printf("Length bVege : %d \n", len(bVege))            // Output: 3
	fmt.Printf("Cap bVege : %d \n", cap(bVege))               // Output: 3

	// append() untuk menambahkan elemen ke slice dengan posisi paling akhir
	fmt.Println("\nappend() usage")
	var cVege = append(vegetables, "kubis")
	for _, cvege := range cVege {
		fmt.Println(cvege)
	}
	// Output: brokoli lobak selada seledri kubis

	// copy() digunakan untuk menyalin isi dari slice maupun array lain
	fmt.Println("\ncopy() usage")
	var newFruits = []string{"apel", "pisang", "ceri"}
	var copiedFruits = make([]string, len(newFruits))
	copy(copiedFruits, newFruits)
	fmt.Println(copiedFruits) // Output: [apel pisang ceri]
	
	// Modifikasi slice asli tidak mempengaruhi slice yang disalin
	newFruits[0] = "aprikot"
	fmt.Println(newFruits)    // Output: [aprikot pisang ceri]
	fmt.Println(copiedFruits) // Output: [apel pisang ceri]

	// Membuat slice dengan make()
	fmt.Println("\nmake() usage")
	var numbers = make([]int, 5, 10)
	fmt.Println(numbers) // Output: [0 0 0 0 0]
	fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))
	// Output: Length: 5, Capacity: 10

	// Pertumbuhan slice dinamis
	fmt.Println("\nDynamic slice growth")
	var animals = []string{"kucing", "anjing"}
	fmt.Printf("Initial slice: %v, Length: %d, Capacity: %d\n", animals, len(animals), cap(animals))
	// Menambahkan elemen ke slice
	animals = append(animals, "gajah", "rubah", "jerapah")
	fmt.Printf("After append: %v, Length: %d, Capacity: %d\n", animals, len(animals), cap(animals))
	// Output: After append: [kucing anjing gajah rubah jerapah], Length: 5, Capacity: 6

	// Menylicing slice
	fmt.Println("\nSlicing slices")
	var colors = []string{"merah", "hijau", "biru", "kuning", "ungu"}
	subset1 := colors[1:3]
	fmt.Println(subset1) // Output: [hijau biru]
	subset2 := colors[:2]
	fmt.Println(subset2) // Output: [merah hijau]
	subset3 := colors[2:]
	fmt.Println(subset3) // Output: [biru kuning ungu]
	subset4 := colors[:]
	fmt.Println(subset4) // Output: [merah hijau biru kuning ungu]

	// Re-slicing
	fmt.Println("\nRe-slicing")
	var numbers2 = []int{1, 2, 3, 4, 5}
	fmt.Println(numbers2) // Output: [1 2 3 4 5]
	narrowView := numbers2[1:4]
	fmt.Println(narrowView) // Output: [2 3 4]
	narrowView = narrowView[:4]
	fmt.Println(narrowView) // Output: [2 3 4 5]

	// Slice multi-dimensional
	fmt.Println("\nMulti-dimensional slices")
	var matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix[0][1]) // Output: 2
	fmt.Println(matrix[2][2]) // Output: 9
	matrix[1][1] = 55
	fmt.Println(matrix)       // Output: [[1 2 3] [4 55 6] [7 8 9]]
}
