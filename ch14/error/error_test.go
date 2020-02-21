package error_test

import(
  "testing"
  "fmt"
  "errors"
)

func someDoing(n int) ([]int, error)  {
  if n <10 || n>100{
    return nil,errors.New("n must rules")
  }
  fiblist:=[]int{1,1}
  for i:=2;i<10;i++{
    fiblist = append(fiblist,fiblist[i-2]+fiblist[i-1])
  }
  return fiblist,nil
}

func TestErrorSomething(t *testing.T)  {
  if v,error:=someDoing(9);error !=nil{
    fmt.Println(error)
  }else{
    fmt.Println("it is right",v)
  }
}
