package singleton_test

import(
  "testing"
  "fmt"
  "sync"
  "unsafe"
)

type Singleton struct{
  
}

var singletInstance *Singleton

var once sync.Once

func onlyOnce() * Singleton {
  once.Do(func ()  {
    fmt.Println("create once")
    singletInstance=new(Singleton)
  })
  return singletInstance
}

func TestSingleOnce(t *testing.T)  {
  var wg sync.WaitGroup
  for i := 0;i < 10;i++ {
    wg.Add(1)
    go func() {
      obj:=onlyOnce()
      wg.Done()
      fmt.Printf("%d\n",unsafe.Pointer(obj))
    }()
  }
  wg.Wait()
}
