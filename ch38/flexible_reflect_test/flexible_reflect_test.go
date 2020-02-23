package flexible_reflect_test

import(
  "testing"
  "fmt"
  "reflect"
  "errors"
)

// func TestIsEqual (t *testing.T)  {
//   a :=map[int]string{1:"one",2:"two",3:"three"}
//   b :=map[int]string{1:"one",2:"two",3:"three",4:"four"}
//   fmt.Println(reflect.DeepEqual(a,b))
//
//   c :=[]int{1,2,3}
//   d :=[]int{1,2,3,}
//   fmt.Println(reflect.DeepEqual(c,d))
// }


type Employee struct{
  EmployeeId   string
  Name         string `format:"normal"`
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

func fillBySettingS(st interface{},settings map[string]interface{}) error  {
  if reflect.TypeOf(st).Kind() != reflect.Ptr{
    return errors.New("the first param  should be a pointer")
  }
  //elem() 获取指针指向的值
  if reflect.TypeOf(st).Elem().Kind() !=reflect.Struct{
    return errors.New("the first param  should be a pointer 2222")
  }

  if settings ==nil{
    return errors.New("settings is null")
  }

  var (
    field reflect.StructField
    ok bool
  )


  for k,v :=range settings {
    if field,ok =(reflect.ValueOf(st)).Elem().Type().FieldByName(k);!ok{
      continue
    }

    if field.Type == reflect.TypeOf(v){
      vstr :=reflect.ValueOf(st)
      vstr =vstr.Elem()
      vstr.FieldByName(k).Set(reflect.ValueOf(v))
    }
  }
  return nil

}


type child struct {
  eating string
  weight float64
}

func TestFillNameAndAge(t *testing.T)  {
  settings :=map[string]interface{}{"Name":"wuqinqiang","Age":30}
  employ :=Employee{}
  if err :=fillBySettingS(&employ,settings);err !=nil{
    t.Fatal(err)
  }




  fmt.Println(employ)

}
