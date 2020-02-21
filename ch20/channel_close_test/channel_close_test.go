package channel_close_test

import(
  "testing"
  "fmt"
  "sync"
)

func dataProducer(ch chan int ,wg *sync.WaitGroup )  {
  go func() {
    for i:=0;i<10;i++{
      ch <- i
    }
    close(ch)
    wg.Done()

  }()
}

func dataReceiver(ch chan int,wg *sync.WaitGroup)  {
  go func() {
    for {
      if res,ok:=<-ch;ok{
        fmt.Println(res)
      }else{
        break
      }
    }
    wg.Done()
  }()
}

//生成消费者
func TestData(t *testing.T)  {
    var wg sync.WaitGroup
    ch:=make(chan int)
    wg.Add(1)
    dataProducer(ch,&wg)
    wg.Add(1)
    dataReceiver(ch,&wg)
    // wg.Add(1)
    // dataReceiver(ch,&wg)
    wg.Wait()
}
