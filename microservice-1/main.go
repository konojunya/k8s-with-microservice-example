package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "hello from microservice 1")
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
	})

	port := ":8081"
	fmt.Println("Server is running on http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
