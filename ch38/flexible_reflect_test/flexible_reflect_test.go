package flexible_reflect_test

import(
  "testing"
  "fmt"
  "reflect"
)

func TestIsEqual (t *testing.T)  {
  a :=map[int]string{1:"one",2:"two",3:"three"}
  b :=map[int]string{1:"one",2:"two",3:"three",4:"four"}
  fmt.Println(reflect.DeepEqual(a,b))

  c :=[]int{1,2,3}
  d :=[]int{1,2,3,}
  fmt.Println(reflect.DeepEqual(c,d))
}
