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

}

func main() {
}
