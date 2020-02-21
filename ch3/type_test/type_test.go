package type_test

import (
  "testing"
  "fmt"
)

//不支持类型隐式转换
func TestImpLicit(t *testing.T)  {
  var a int64=1
  var b int64=2
  a,b=b,a
  fmt.Println(a,b)
}

//不能进行指针运算
func TestPoint(t *testing.T)  {
  a:=1
  aPtr:=&a
  //aPtr=aPtr+1 //不支持
  fmt.Println(a,aPtr)
  t.Log(a,aPtr)
}

func TestString(t *testing.T)  {
  var s string
  fmt.Println("前面"+s)
  fmt.Println(len(s))
  if s==""{
    fmt.Println("空了不好")
  }
}
