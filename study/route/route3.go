package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//路由分组|中间件

func listPosts(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"文章列表")
}

func createPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "发布文章")
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "修改文章")
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "删除文章")
}

func showPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "文章详情")
}

//中间件
func checkToken(next http.Handler) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		token:=r.FormValue("token")
		if token=="curry.com"{
			log.Printf("token is success:%s\n",r.RequestURI)
			next.ServeHTTP(w,r)
		}else {
			http.Error(w,"Forbidden",http.StatusForbidden)
		}
	})
}


func main()  {
	r:=mux.NewRouter()

	postRouter:=r.PathPrefix("/posts").Subrouter()
	postRouter.Use(checkToken)
	postRouter.HandleFunc("/",listPosts).Methods("GET")
	postRouter.HandleFunc("/create", createPost).Methods("POST")
	postRouter.HandleFunc("/update", updatePost).Methods("PUT")
	postRouter.HandleFunc("/delete", deletePost).Methods("DELETE")
	postRouter.HandleFunc("/show", showPost).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080",r))
}
