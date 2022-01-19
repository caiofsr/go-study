package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greets(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func HandlerMyGreetings(w http.ResponseWriter, r *http.Request) {
	Greets(w, "Caio")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(HandlerMyGreetings))

	if err != nil {
		fmt.Println(err)
	}
}
