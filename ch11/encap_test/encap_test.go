package encap_test

import(
  "testing"
  "fmt"
  "unsafe"

)

type Employee struct{
  Id string
  Name string
  Age int
}

//使用引用的时候不需要重新开辟空间，同一个地址，节省开销
func (e *Employee) String() string {
  fmt.Printf("address is %x\n",unsafe.Pointer(&e.Name))
  return fmt.Sprintf("Id:%s/Name:%s/Age:%d",e.Id,e.Name,e.Age)
}

func TestCreateEmploy(t *testing.T)  {
e:=Employee{"0","Bob",20}
fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))
//fmt.Println(e)
//e1:=Employee{Name:"wuqinqiang",Age:16}
//fmt.Println(e1.Name)
fmt.Println(e.String())

}
