package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type TemplateData struct {
	Title    string
	Datetime string
	Unixtime string
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	// 親テンプレート(index.html)と子テンプレート(body.html)
	tmpl := template.Must(template.ParseFiles("views/index.html", "views/body.html"))

	title := "現在の時刻"

	// bodyに表示するコンテンツ
	// 時刻オブジェクトを文字列に変換
	datetime := fmt.Sprint(time.Now())
	unixtime := fmt.Sprint(time.Now().Unix())

	// テンプレートを実行して出力
	templatedate := TemplateData{title, datetime, unixtime}
	if err := tmpl.ExecuteTemplate(w, "base", templatedate); err != nil {
		fmt.Println(err)
	}
}

func main() {
	// http.HandleFuncにルーティングと処理する関数を登録
	http.HandleFunc("/", HelloServer)

	// ログ出力
	log.Printf("Start Go HTTP Server")

	// http.ListenAndServeで待ち受けるportを指定
	err := http.ListenAndServe(":4001", nil)

	// エラー処理
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
