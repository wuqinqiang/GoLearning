package share_mem_test

import(
  "testing"
  "fmt"
  "time"
  "sync"
)

//非线程安全
func TestShare(t *testing.T)  {
  count :=0
  for i:=0;i<5000;i++{
    go func() {
      count++
    }()
  }
  time.Sleep(1*time.Second)
  fmt.Println(count)
}

//锁  还是会错误，如果把sleep去掉
func TestShareSafe(t *testing.T)  {
  var mut sync.Mutex
  count :=0
  for i:=0;i<5000;i++{
    go func() {
      defer func ()  {
        mut.Unlock()
      }()
      mut.Lock()
      count++
    }()
  }
  //time.Sleep(1*time.Second)
  fmt.Println(count)
}

//waitGroup
func TestShareWaitGroup(t *testing.T)  {
  var mut sync.Mutex
  var wg sync.WaitGroup
  count :=0
  for i:=0;i<5000;i++{
    wg.Add(1)
    go func() {
      defer func ()  {
        mut.Unlock()
      }()
      mut.Lock()
      count++
      wg.Done()
    }()
  }
  wg.Wait()
  fmt.Println(count)
}
