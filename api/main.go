package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Go API running on port 4000")
	http.ListenAndServe(":4000", nil)
}
