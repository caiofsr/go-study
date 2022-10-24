package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum =>", a+b)
	fmt.Println("Subtraction =>", a-b)
	fmt.Println("Div =>", a/b)
	fmt.Println("Multiplication =>", a*b)
	fmt.Println("Module =>", a%b)

	//bitwise
	fmt.Println("And =>", a&b) // 11 & 10 = 10
	fmt.Println("Or =>", a|b) // 11 | 10 = 11
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	fmt.Println("Greater =>", math.Max(float64(a), float64(b)))
	fmt.Println("Grater =>", math.Max(c, d))
	fmt.Println("Exponencial =>", math.Pow(c, d))

}