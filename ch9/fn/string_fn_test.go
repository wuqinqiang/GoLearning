package chat_test

import(
  "testing"
  "fmt"
  "strings"
  "strconv"
)

//字符串切片
func TestStringFn(t *testing.T)  {
  s:="A,B,C"
  parts:=strings.Split(s,",")
  for _,part:=range parts{
    fmt.Println(part)
  }
  fmt.Println(strings.Join(parts,"-"))
}

//字符串转整形 整形转字符串 字符串转整形是有可能出现错误的，所以需要判断
func TestConv(t *testing.T)  {
  s:=strconv.Itoa(10)
  fmt.Println("hello"+s)
  if i,error:=strconv.Atoi("10");error==nil{
    fmt.Println(5+i)
  }
}
