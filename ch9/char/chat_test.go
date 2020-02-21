package chat_test

import(
  "testing"
  "fmt"
)

func TestString(t *testing.T)  {
var s string
fmt.Println(s)
s="hello"
fmt.Println(len(s))
s="\xE4\xB8\xA5"
fmt.Println(s)
}
