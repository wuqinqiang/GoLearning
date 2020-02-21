package switch_test

import (
  "testing"
  "fmt"
)

func TestSwitch(t * testing.T)  {
a:=11;
switch{
case a>=0 && a<=5:
  fmt.Println("a<=5")
case a>=10:
  fmt.Println("a>=10")
default:
  fmt.Println("no")
}

}
