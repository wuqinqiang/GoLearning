package main

import (
	"fmt"
	"net/http"
)

type Hello struct {
}

type Curry struct {

}

func sayHello(w http.ResponseWriter, r *http.Request) {

}



func (handler *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"hello Golang")
}

func (curry *Curry) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	helloCurry(w)
}

func helloCurry(w http.ResponseWriter)  {
	fmt.Fprintf(w,"hello curry")
}

func main() {
	hello := Hello{}
	curry:=Curry{}
	serve:=http.Server{
		Addr:              ":9091",
	}
	http.Handle("/hello",&hello)
	http.Handle("/curry",&curry)
	serve.ListenAndServe()
}
