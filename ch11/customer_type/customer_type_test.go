package customer_type

import(
  "testing"
  "fmt"
  "time"
)

var IntConv func (op int) int

func timeSpent(inner IntConv) IntConv  {
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
  tsSf:=timeSpent(slowFun)
  fmt.Println(tsSf(9))
}
