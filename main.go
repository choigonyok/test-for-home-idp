package main

import (
	"fmt"
	"net/http"

	"github.com/google/go-github/v63/github"
)

// 간단한 핸들러 함수
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	github.NewClient(http.DefaultClient)
	// "/hello" 경로에 대한 요청을 helloHandler로 	처리
	http.HandleFunc("/hello", helloHandler)

	// 서버 시작
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
