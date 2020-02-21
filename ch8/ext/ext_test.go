package ext_test

import(
  "testing"
  "fmt"
)
func TestMapWithFunValue(t *testing.T)  {
m:=map[int]func (op int) int {}
m[1]=func (op int) int  {
  return op
}
m[2]=func (op int) int  {
  return op * op
}
m[3]=func (op int) int  {
  return op * op * op
}
fmt.Println(m[1](2),m[2](2),m[3](2))
}

//set
func TestMapSetExist(t *testing.T)  {
  set:=map[int]bool{}
  set[1]=true
  if set[1]{
    fmt.Println("key 1 exist")
  }else {
    fmt.Println("key 1 no exist")
  }
  delete(set,1)
  if set[1]{
    fmt.Println("key 1 exist")
  }else {
    fmt.Println("key 1 no exist")
  }

}
