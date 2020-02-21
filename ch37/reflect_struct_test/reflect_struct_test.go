package reflect_struct_test

import(
  "testing"
  "fmt"
  "reflect"
)

type Employee struct{
  EmployeeId   string
  Name         string //`format:"normal"`
  Age          int
}

func (e *Employee) UpdateAge(newAge int)  {
  e.Age=newAge
}

type Customer struct {
  CookieId string
  Name     string
  Age      int
}

//反射
func TestInvokeByName(t *testing.T)  {
  e:=&Employee{"1","wuqinqiang",24}
  fmt.Println("Name:",reflect.ValueOf(*e).FieldByName("Name"))
  // if name,ok:=reflect.TypeOf(*e).FieldByName("Name");!ok{
  //   fmt.Println("no name")
  // }else{
  //   fmt.Println("name:",name)
  // }
reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(5)})
fmt.Println(e)

}
