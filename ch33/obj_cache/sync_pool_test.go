package sync_pool_test

import(
  "testing"
  "fmt"
  "sync"
//  "runtime"
)

func TestSyncPool(t *testing.T)  {
  pool:=&sync.Pool{
    New:func () interface{}  {
      fmt.Println("first create")
      return 100
    },
  }
  v:=pool.Get().(int)
  fmt.Println(v)
  pool.Put(3)
  //runtime.GC()
  v1,_:=pool.Get().(int)
  fmt.Println(v1)
}
