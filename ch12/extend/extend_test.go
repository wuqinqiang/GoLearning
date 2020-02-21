package extend_test
import (
  "testing"
  "fmt"
)

type Code string

type Programmer interface {
  writeHelloWorld() Code
}

type GoProgrammer struct {

}

func (g *GoProgrammer) writeHelloWorld() Code  {
  return "go hello";
}

type JavaProgrammer struct {

}

func (j *JavaProgrammer) writeHelloWorld() Code {
  return "java hello";
}

func writeFirt(p Programmer)  {
  fmt.Println(p,p.writeHelloWorld())
}


func TestProgrammer(t *testing.T)  {
  goP:=new(GoProgrammer)
  javaP:=new(JavaProgrammer)
  writeFirt(goP)
  writeFirt(javaP)


}
