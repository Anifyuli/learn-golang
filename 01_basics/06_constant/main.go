package main

import (
	"fmt"
	"math"
)

func main() {
	// Cara menuliskan konstanta tanpa menyebutkan tipedata
	const pi = 3.14

	// Di Golang, kita dapat menuliskan konstanta lebih dari satu secara bersamaan tanpa mengulang kata kunci "const"
	const (
		planet string = "Earth"
		gravityConst float64 = 6.67e11
	)

	// Menghitung luas lingkaran
	result := pi * 7 * 7
	fmt.Printf("Result is : %f cm \n", result)

	// Menghitung nilai gravitasi
	mOne, mTwo := 2, 4
	radius := 21
	force := gravityConst * float64(mOne) * float64(mTwo) / math.Pow(float64(radius), 2)
	fmt.Printf("Force of %s is : %f N \n", planet, force)
}