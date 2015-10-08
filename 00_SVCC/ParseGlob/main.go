package main

import (
	"html/template"
	"log"
	"os"
	"fmt"
)

type Page struct  {
	Title string
	Body string

}

func main() {

	tpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, Page{
		Title: "Which file?",
		Body: "hello page 1",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n****************")

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.html", Page{
		Title: "specifying tpl.html",
		Body: "hello page 1",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n****************")

	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.html", Page{
		Title: "specifiying tpl2.gohtml",
		Body: "hello page 2",
	})
	if err != nil {
		log.Fatalln(err)
	}

}