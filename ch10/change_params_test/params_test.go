package chat_test

import(
  "testing"
  "fmt"
)

//可变成参数
func Sum(agrs ...int)int  {
  res:=0
  for _,val:=range agrs{
    res +=val
  }
  return res;
}

func TestChangeParams(t *testing.T)  {
  a:=Sum(1,2,3,4,5)
  b:=Sum(1,2,3,4,5,6)
  fmt.Println(a,b)
}

func clear()  {
  fmt.Println("clear anything")
}

//延迟函数 defer
func TestDeferFun(t *testing.T)  {
  defer clear()
  fmt.Println("先走这里")
//panic("error")
  //fmt.Println("最后")


}
