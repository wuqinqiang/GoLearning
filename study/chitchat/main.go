package main

import (
	. "github.com/wuqinqiang/chitchat/routes"
	"log"
	"net/http"
)

func startWebServer(port string) {
	r := NewRoute()

	//处理静态文件
	assets:=http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",assets))

	http.Handle("/", r) //通过 route.go 中定义的路由器来分发路由请求
	log.Println("starting HTTP service at" + port)
	err := http.ListenAndServe(":"+port, nil) //启动协程监听请求
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}

}

func main() {
	startWebServer("8080")
}
