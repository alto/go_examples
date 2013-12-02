package main

import (
  "log"
  // "fmt"
  "net/http"
  "html/template"
)

type Page struct {
    Title string
}

func renderPage(w http.ResponseWriter, r *http.Request) {
  p := Page{Title: "The Title"}
  t, _ := template.ParseFiles("page.html")
  t.Execute(w, p)
}

func main() {
  http.HandleFunc("/", renderPage)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
