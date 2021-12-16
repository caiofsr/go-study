package main

import "fmt"

const portuguese = "portuguese"
const french = "french"
const spanish = "spanish"
const prefixHola = "Hola, "
const prefixHello = "Hello, "
const prefixOla = "Olá, "
const prefixBonjour = "Bonjour, "

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingsPrefix(language) + name
}

func greetingsPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = prefixOla

	case french:
		prefix = prefixBonjour

	case spanish:
		prefix = prefixHola

	default:
		prefix = prefixHello
	}

	return
}

func main() {
	fmt.Println(Hello("mundo", ""))
}
