package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//获取url的绝对路径
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//在url路径上添加静态资源目录
	path = filepath.Join(h.staticPath, path)

	//检查对应资源文件是否存在
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		//文件不存在返回入口html文档内容作为相应
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))

		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//一切顺利
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	router := mux.NewRouter()
	spa := spaHandler{
		staticPath: "dist",
		indexPath:  "index.html",
	}
	router.PathPrefix("/").Handler(spa)
	srv := &http.Server{
		Addr:         "127.0.0.1:8000",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
