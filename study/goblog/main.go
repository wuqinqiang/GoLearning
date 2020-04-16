package main

import (
	. "github.com/wuqinqiang/goblog/goblog/routes"
	"log"
	"net/http"
)

func main() {
	startWebServe("8080")
}

func startWebServe(port string) {
	r := NewRouter()
	http.Handle("/", r)
	log.Printf("start HTTP service at" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
