package cancel_test
import(
  "testing"
  "fmt"
  "time"
)

func isCanceld(cancelChan chan struct{}) bool  {
  select{
  case <-cancelChan:
    return true
  default:
    return false
  }
}

func cacnel_1(cancelChan chan struct{})  {
  cancelChan <- struct{}{}
}

func cacnel_2(cancelChan chan struct{})  {
  close(cancelChan)
}

func TestCancel(t *testing.T)  {
  cancelChan:=make(chan struct{},0)
  for i := 0; i < 5; i++ {
    go func(i int,cancelCh chan struct{}) {
      for{
        if isCanceld(cancelCh){
          break
        }
        time.Sleep(time.Millisecond*5)
      }

      fmt.Println(i,"Cancelled")
    }(i,cancelChan)
  }
  cacnel_2(cancelChan)
  time.Sleep(time.Second*1)
}
