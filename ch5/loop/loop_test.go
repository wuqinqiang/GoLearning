package loop_test

import(
  "testing"
  "fmt"
)

func TestWhileLoop(t *testing.T)  {
  n:=0;
  for n<=5{
    fmt.Println(n)
    n++
  }
}
