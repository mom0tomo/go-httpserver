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
	// parent template and child template
	tmpl := template.Must(template.ParseFiles("views/index.html", "views/body.html"))

	title := "現在の時刻"

	// contents on body
	// convert time to string
	datetime := fmt.Sprint(time.Now())
	unixtime := fmt.Sprint(time.Now().Unix())

	// output the template
	templatedate := TemplateData{title, datetime, unixtime}
	if err := tmpl.ExecuteTemplate(w, "base", templatedate); err != nil {
		fmt.Println(err)
	}
}

func main() {
}
