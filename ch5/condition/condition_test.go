package condition_test

import(
  "testing"
  "fmt"
)
//go 的if语句  a必须是布尔型  
func TestIfMultiSec(t *testing.T)  {
  if a:=12==1;a{
    fmt.Println("1==1")
  }else{
    fmt.Println("1 !=1")
  }
}
