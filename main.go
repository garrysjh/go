package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, World!")
		fmt.Println(fmt.Sprintf("Wrote %d bytes", n))
		if err == nil {
			fmt.Println("No error")
		} else {
			fmt.Println(fmt.Sprintf("Error: %s", err))
		}
	})

	_ = http.ListenAndServe(":8080", nil)
}
