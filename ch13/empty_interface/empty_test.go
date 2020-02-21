package empty_test

import (
  "testing"
  "fmt"
)

func doSometing(p interface{})  {
  if i,ok:=p.(int);ok{
    fmt.Println("Int",i)
    return
  }

  if i,ok:=p.(string);ok{
    fmt.Println("string",i)
    return
  }

  fmt.Println("it is no")

}

func TestEmpty(t *testing.T)  {
  doSometing("11")
}
