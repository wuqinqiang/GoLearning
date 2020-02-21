package context_test
import(
  "testing"
  "fmt"
  "time"
  "context"
)

func isCanceld(ctx context.Context) bool  {
  select{
  case <-ctx.Done():
    return true
  default:
    return false
  }
}


func TestCancel(t *testing.T)  {
  ctx,cancel:=context.WithCancel(context.Background())
  for i := 0; i < 5; i++ {
    go func(i int,ctx context.Context) {
      for{
        if isCanceld(ctx){
          break
        }
        time.Sleep(time.Millisecond*5)
      }

      fmt.Println(i,"Cancelled")
    }(i,ctx)
  }
  cancel()
  time.Sleep(time.Second*1)
}
