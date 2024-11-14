package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/contacto",Contact)

    server:=http.ListenAndServe(":8080", router)
    log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w,"<h1>Pagina index de mi sitio Web</h1>")
}

func Contact(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w,"<h1>Contacto</h1>")
}