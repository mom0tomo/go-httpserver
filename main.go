package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, 世界\n")
}

func TestServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TEST\n")
}

func main() {
	// http.HandleFuncにルーティングと処理する関数を登録
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/test", TestServer)

	// ログ出力
	log.Printf("Start Go HTTP Server")

	// http.ListenAndServeで待ち受けるportを指定
	err := http.ListenAndServe(":4000", nil)

	// エラー処理
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
