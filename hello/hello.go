package main

import "fmt"

func Hello(name string) string {
	const prefixHello = "Hello, "

	if name == "" {
		name = "World"
	}

	return prefixHello + name
}

func main() {
	fmt.Println(Hello("mundo"))
}
