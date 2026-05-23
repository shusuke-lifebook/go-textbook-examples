package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Go!")
	})

	fmt.Println("サーバー起動: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Fprintf(os.Stderr, "サーバーエラー: %v\n", err)
		os.Exit(1)
	}
}
