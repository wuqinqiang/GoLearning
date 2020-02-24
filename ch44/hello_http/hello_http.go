package main

import (
  "fmt"
  "net/http"
  "time"
)

func main()  {
  http.HandleFunc("/",func (w http.ResponseWriter,r *http.Request)  {
    fmt.Fprintf(w,"hello world!")
  })
  http.HandleFunc("/time",func (w http.ResponseWriter,r *http.Request)  {
    t := time.Now()
    timeStr :=fmt.Sprintf("time is :%s",t)
    w.Write([]byte(timeStr))
  })
  http.ListenAndServe(":8080",nil)
}
