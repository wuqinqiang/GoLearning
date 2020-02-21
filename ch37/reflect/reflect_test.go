package reflect_test

import(
  "testing"
  "fmt"
  "reflect"
)

func checkType(v interface{})  {
  res:=reflect.TypeOf(v)
  switch res.Kind() {
  case reflect.Float32,reflect.Float64 :
    fmt.Println("float")
  case reflect.Int,reflect.Int32,reflect.Int64:
    fmt.Println("int")
  default:
    fmt.Println("unKnown",res)
  }
}


func TestTypeAndValue(t *testing.T)  {
  var  number int =10
  fmt.Println(reflect.TypeOf(number),reflect.ValueOf(number))
}

func TestWhichType(t *testing.T)  {
  var f int =23
  checkType(f)
}
