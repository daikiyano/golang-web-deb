package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout,"tpl.html",`this is valiable test.`)
	if err != nil {
		log.Fatalln(err)
	}
}
