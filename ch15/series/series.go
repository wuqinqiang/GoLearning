package series

import "fmt"
func init()  {
  fmt.Println("init1")
}

func init()  {
  fmt.Println("init2")
}

func SomeDoing(n int) []int  {
  fiblist:=[]int{1,1}
  for i:=2;i<10;i++{
    fiblist = append(fiblist,fiblist[i-2]+fiblist[i-1])
  }
  return fiblist
}
