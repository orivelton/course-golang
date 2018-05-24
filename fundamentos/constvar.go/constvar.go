package main

import (
	"math"
	"fmt"
)

func main() {
	const PI float64 = 3.1415 // float64 tipo padrão de um literal
	var raio = 3.2

	//Sugar sintax
	//area = // = atribuindo um valor

	area := PI * math.Pow(raio, 2) // := quando a var ainda não existe

	fmt.Println("A área da circuferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "Epa!"

	fmt.Println(g, h, i)
}