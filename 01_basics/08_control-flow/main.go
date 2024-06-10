package main

import "fmt"

func main() {
	// Penerapan if-else
	isLove := true

	if isLove {
		fmt.Println("I love her")
	} else {
		fmt.Println("I'm not love her")
	}

	// Penerapan if-else dengan variabel sementara
	maxLoveLevel := 100.0

	if percent := 70.0 / maxLoveLevel; percent >= 0.3*maxLoveLevel {
		fmt.Println("You not seriously love her")
	} else {
		fmt.Println("You really love her")
	}

	// Penerapan switch-case
	var point = 2

	switch point {
	case 8:
		fmt.Println("Perfect")
	case 7, 6, 5, 4:
		fmt.Println("Awesome")
		fallthrough // Digunakan untuk memaksa untuk menjalankan proses case selanjutnya
	default:
		{
			fmt.Println("Not bad")
			fmt.Println("You can be better!")
		}
	}
}