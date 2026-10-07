package main

import (
	"fmt"
)

func main() {

	// ini program untuk mendefined variable
	nama := "budi"
	umur := 25
	tinggi := 165

	const phi = 3.14

	fmt.Println(nama, umur, tinggi, phi)

	if umur >= 17 {
		fmt.Printf("%s sudah berumur \n", nama)
	} else {
		fmt.Printf("%s sudah berumur \n", nama)
	}

	n := 5
	for i := 1; i <= 5; i++ {
		fmt.Println(n)

		n += 5
	}
	// for i := 1; i <= 5; i++ {
	// 	n *= 5
	// 	fmt.Println(n)
	// }

}
