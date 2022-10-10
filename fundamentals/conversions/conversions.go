package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	grade := 6.9
	finalGrade := int(grade)
	fmt.Println(finalGrade)

	fmt.Println("teste " + string(97))
	fmt.Println("teste " + strconv.Itoa(123))
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("True")
	}
}
