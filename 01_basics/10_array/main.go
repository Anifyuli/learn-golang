package main

import "fmt"

func main() {
	/*
		Cara membuat array dalam Golang dengan mendefinisikan nama variabel dan tipedata membernya
		diikuti dengan mendaftarkan setiap member beserta indeksnya
	*/
	var waifu [4]string
	waifu[0] = "Marin Kitagawa"
	waifu[1] = "Sylphiette"
	waifu[2] = "Mahiru Shina"
	waifu[3] = "Alise Lendrott"

	fmt.Println("Waifu List : ")

	// Menampilkan isi array dengan for
	for i := 0; i < len(waifu); i++ {
		fmt.Printf("- %s \n", waifu[i])
	}

	fmt.Printf("\n")

	// Array juga dapat didefinisikan seperti berikut
	var husbando = [4]string{"Zagan", "Flio", "Lloyd Forger", "Kraft Laurence"}

	fmt.Println("Husbando List : ")
	// len() dipakai untuk menghitung panjang atau jumlah member dari sebuah array
	for i := 0; i < len(husbando); i++ {
		fmt.Printf("- %s \n", husbando[i])
	}

	// Jumlah member array dapat diganti dengan elipsis (...) jika tidak ditentukan panjang datanya
	var animeTitles = [...]string{"Kimetsu No Yaiba", "Jujutsu Kaisen", "Tensura", "Konosuba", "Wind Breaker"}

	fmt.Printf("\n")

	fmt.Println("Anime Titles List : ")
	for _, animeTitle := range animeTitles {
		fmt.Printf("- %s \n", animeTitle)
	}

	var demonFruits = make([]string, 3)
	demonFruits[0] = "Gomu Gomu no Mi"
	demonFruits[1] = "Bara Bara no Mi"
	demonFruits[2] = "Sube Sube no Mi"

	fmt.Printf("\n")

	fmt.Println("Demon Fruit List : ")
	for _, demonFruit := range demonFruits {
		fmt.Printf("- %s \n", demonFruit)
	}

}
