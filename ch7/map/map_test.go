package map_test

import(
  "testing"
  "fmt"
)

//创建map
func TestInitMap(t *testing.T)  {
  m1:=map[int]int{1:1,2:2,3:3}
  fmt.Println("m1长度是",len(m1))
  fmt.Println(m1[2])
  m2:=map[int]int{}
  m2[2]=10
  fmt.Println("m2的长度是",len(m2))
  m3:=make(map[int]int, 10)
  fmt.Println("m3的长度是",len(m3))
}

//map访问不存在的键结果为0，区分是否存在和存在值为0的情况

func TestExistsKey(t *testing.T)  {
  m3:=map[int]int{};
  fmt.Println(m3[2])
  m3[2]=0
  fmt.Println(m3[4])
  if value,ok:=m3[2];ok{
    fmt.Println("key 2 exist，value is ",value)
  }else{
    fmt.Println("key 2 dont exist")
  }
}

 //map遍历
func TestMapTravel(t *testing.T)  {
  m1:=map[int]int{1:1,2:2,3:3}
  for index,value:=range m1{
    fmt.Println(index,value)
  }
}
