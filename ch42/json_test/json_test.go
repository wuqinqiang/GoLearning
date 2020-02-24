package jsontest

import (
  "encoding/json"
  "fmt"
  "testing"
)

var jsonStr = `{
  "basic_info":{
    "name":"Mike",
    "age":30
  },
  "job_info":{
    "skills":["Java","Go","C"]
  }
  }    `

//这种效率低下 是用了反射的机制
  func TestJson(t *testing.T)  {
    e :=new(Employee)
    //将json 字符串解码到相应的数据结构(e)
    err :=json.Unmarshal([]byte(jsonStr),e)
    if err !=nil{
      t.Error(err)
    }
    fmt.Println(*e)
    //编码
    if v,err :=json.Marshal(e);err == nil{
      fmt.Println(string(v))
    }else{
      t.Error(err)
    }
  }
