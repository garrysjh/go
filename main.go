package main

import "fmt"

func main() {


	http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf("Hello world!")
}
}
