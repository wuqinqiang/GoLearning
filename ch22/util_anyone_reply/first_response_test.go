package first_response_test

import(
  "testing"
  "fmt"
  "time"
  "runtime"
)

func runTask(i int) string  {
    time.Sleep(time.Millisecond*10)
  return fmt.Sprintf("the result is from %d",i)
}

//  ch:=make(chan string,10)

//上面这句话是重点，使用buffer channel 防止协程泄漏，一般情况下，如果不设置buffer的话
//那么如果没有消费者的话 生产者是阻塞的状态，如果有buffer 那么生产者只顾把消息发送到 buffer即可
func FirstResponse() string  {
  ch:=make(chan string,10)
  for i := 0;i < 10;i++ {
    go func(i int) {
      res:=runTask(i)
      ch<-res
    }(i)
  }

  return <-ch
}


func TestFirst(t *testing.T)  {
  fmt.Println("before:",runtime.NumGoroutine())
  fmt.Println(FirstResponse())
  time.Sleep(time.Second*1)
  fmt.Println("after:",runtime.NumGoroutine())

}
