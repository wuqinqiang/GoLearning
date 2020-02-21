package constant_test

import(
  "testing"
  "fmt"
  )

const(
  Monday=iota+1
  Tuesday
  Wednesday
)

const(
  Readable = 1<<iota
  Writable
  Executablle
)

func TestConstantTry(t * testing.T)  {

  fmt.Println(Monday,Wednesday)
}

func TestConStantTry1(t *testing.T)  {
  a:=5 //1=0001 2=0010 3 =0011 4=0100 5= 0101 6=0110
  fmt.Println(a&Readable == Readable,a&Writable==Writable,a&Executablle==Executablle);
}
