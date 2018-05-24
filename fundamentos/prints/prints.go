package main

import (
	"fmt"
)

func main () {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println("Mesma")
	fmt.Println(" linha.")

	x := 3.141516

	//fmt.Println("O valor de x é " + x)

	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é " + xs)

	fmt.Println("O valor de x é", x)
}