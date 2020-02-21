package chat_test

import(
  "testing"
  "fmt"
  "math/rand"
  "time"
)

func returnTwoValue()(int,int)  {
  return rand.Intn(10),rand.Intn(20)
}

func timeSpent(inner func(op int) int) func (op int) int  {
  return func (n int) int  {
    start:=time.Now()
    ret:=inner(n)
    fmt.Println("time spent:",time.Since(start).Seconds())
    return ret
  }
}

func slowFun(op int) int  {
  time.Sleep(time.Second*1)
  return op

}

func TestFnReturn(t *testing.T)  {
  //_,b:=returnTwoValue()
  tsSf:=timeSpent(slowFun)
  fmt.Println(tsSf(9))
}
