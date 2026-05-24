package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", helloHandler)
	mux.HandleFunc("GET /search", searchHandler)
	mux.HandleFunc("GET /users/{id}", getUser)
	mux.HandleFunc("POST /users", createUser)

	// 全ルートにミドルウェアを適用
	handler := loggingMiddleware(mux)

	fmt.Println("サーバー起動: http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		fmt.Fprintf(os.Stderr, "サーバーエラー: %v\n", err)
		os.Exit(1)
	}
}
