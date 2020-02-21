package client

import(
  "testing"
  "ch15/series"
  "fmt"
)

func TestClient(t *testing.T)  {
  res:=series.SomeDoing(10)
  fmt.Println(res)
}
