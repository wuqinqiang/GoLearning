package async_service_test

import(
  "testing"
  "fmt"
  "time"
)

func service() string {
  time.Sleep(1*time.Second)
  return " done"
}

func otherService()  {
  fmt.Println("1")
  time.Sleep(2*time.Second)
  fmt.Println("2")
}

func AsynService() chan string {
  retch:=make(chan string,1)
  go func() {
    ret:=service()
    fmt.Println("3")
    retch <- ret
    fmt.Println("4")
  }()
  return retch
}

// func TestService(t *testing.T)  {
//   fmt.Println(service())
//   otherService()
// }

func TestAsynService(t *testing.T)  {
  retch:=AsynService()
  otherService()
  fmt.Println(<-retch)
  time.Sleep(time.Second*1)

}
