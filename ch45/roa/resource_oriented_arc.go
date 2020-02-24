package main

import(
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "github.com/julienschmidt/httprouter"

)

type Employee struct {
  ID    string     `json:"id"`
  Name  string     `json:"name"`
  Age   int        `json:"age"`
}

var employeeDB map[string]*Employee

func init()  {
  employeeDB = map[string]*Employee{}
  employeeDB["Mike"]=&Employee{"e-1","Mike",35}
  employeeDB["Nick"]=&Employee{"e-2","Nick",20}
}

func Index(w http.ResponseWriter,r *http.Request, _ httprouter.Params)  {
  fmt.Fprint(w,"welcome!\n")
}

func GetEmployeeByName(w http.ResponseWriter,r *http.Request, ps httprouter.Params)  {
  qName :=ps.ByName("name")
  var (
    ok        bool
    info      *Employee
    infoJson  []byte
    err       error
  )
  if info,ok = employeeDB[qName];!ok{
    w.Write([]byte("not found"))
  }

  if infoJson,err = json.Marshal(info);err !=nil{
    w.Write([]byte(fmt.Sprintf("{\"error\":,\"%s\"}",err)))
    return
  }

  w.Write(infoJson)

}

func main()  {
  route :=httprouter.New()
  route.GET("/",Index)
  route.GET("/employee/:name",GetEmployeeByName)
  log.Fatal(http.ListenAndServe(":8080",route))
}
