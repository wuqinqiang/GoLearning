package extension_test

import (
  "testing"
  "fmt"
)

type Pet struct {

}

func (p *Pet) Speak()  {
  fmt.Println("...")
}

func (p * Pet) SpeakTo(host string)  {
  p.Speak();
  fmt.Println("你好啊",host)
}

type Dog struct {
  Pet
}

// func (d *Dog) Speak()  {
//   d.p.Speak()
// }
func (d *Dog) Speak()  {
  fmt.Println("Wang!")
}
//
// func (d * Dog) SpeakTo(host string)  {
//   d.p.SpeakTo(host)
// }

//扩展和复用
func TestExension(t *testing.T)  {
  dog:=new(Dog)
  dog.SpeakTo("wuqin")
  dog.Speak()
}
