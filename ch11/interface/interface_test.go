package interface_test

import(
  "testing"
  "fmt"
)

type Progammer interface {
  WriteHelloWorld() string
}

type GoProgammer struct{

}

func (p *GoProgammer) WriteHelloWorld() string  {
  return "fmt.Printf(\"hello world\")"
}

func TestProgammer(t *testing.T)  {
  var p Progammer
  p=new(GoProgammer)
  fmt.Println(p.WriteHelloWorld())
}
