package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Curry1 struct {

}

func (curry *Curry1) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	params:=mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"hello:%s",params["name"])
}

//使用了gorilla/mux 包来实现路由

func sayHello(w http.ResponseWriter,r *http.Request)  {
	params:=mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"hello world:%s",params["name"])
}

func main()  {
	r:=mux.NewRouter()
	r.HandleFunc("/hello/{name:[a-z]+}",sayHello) //正则匹配
	//r.Handle("/zh/{name}",&Curry1{}).Methods("GET").Host("goweb.test")
	r.Handle("/zh/{name}",&Curry1{}).Methods("GET").Schemes("https")

	log.Fatal(http.ListenAndServe(":8080",r))
}
