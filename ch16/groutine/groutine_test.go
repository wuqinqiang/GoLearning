package groutine_test

import(
  "testing"
  "fmt"
  //"time"
)

//协程
func TestGroutineExplame(t *testing.T)  {
  for i:=0;i<10;i++{
    go func (i int)  {
      fmt.Println(i)
    }(i)
  }

//  time.Sleep(time.Millisecond * 50)
}
