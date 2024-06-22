package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/janhavialekar/mongoapi/router"
)

func main() {
	fmt.Println("hello i'm in go")
	r := router.Router()
	fmt.Println("Server is starting")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")
}
