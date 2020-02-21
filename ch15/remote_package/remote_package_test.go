package remote_test

import (
  "testing"
  cm "github.com/easierway/concurrent_map"

  "fmt"
)

func TestRemote(t *testing.T)  {
  m:=cm.CreateConcurrentMap(99)
  m.Set(cm.StrKey("key"),10)
  fmt.Println(m.Get(cm.StrKey("key")))
}
