package main

import "fmt"

func main() {
	// Menulis variabel dengan menyebutkan tipedata
	var firstName string = "Jhon"
	var lastName string = "Doe"

	// Menulis variabel tanpa menyebutkan tipedata
	health := "good"

	// Multi variabel
	var one, two, three string = "ich", "ni", "san" 
	four, five, six := "yon", "go", "roku"

	// Reserved variable atau variabel yang nilainya tidak dipakai
	_ = 1;
	_ = false;

	// Kata kunci "new" yang dipakai untuk mendefinisikan variabel baru beserta letaknya dalam memori
	eleven := new(int)
	*eleven = 11

	fmt.Printf("Good morning, %s %s \n", firstName, lastName)
	fmt.Printf("Him health is %s \n", health)
	fmt.Printf("One is %s, two is %s, three is %s \n", one, two, three)
	fmt.Printf("Four is %s, five is %s, six is %s \n", four, five, six)

	// Jika langsung dicetak biasa, maka hanya memunculkan letaknya dalam memori
	fmt.Println(eleven)
	
	/*	Jika ingin menampilkan nilai dari variabel yang didefinisikan dengan "new", 
		tambahkan asterisk (*) didepan nama variabel untuk deference variabel tersebut
	*/
	fmt.Println(*eleven)
}