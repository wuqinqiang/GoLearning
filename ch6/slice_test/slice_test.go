package slice_test
import(
  "testing"
  "fmt"
)

//数组
func TestArrayInit(t *testing.T)  {
  var s0 []int

  fmt.Println(len(s0),cap(s0))

  s0=append(s0,1)
  fmt.Println(len(s0),cap(s0))


//使用make创建切片，3表示长度 5表示容量
  s1:=make([]int, 3,5)
  fmt.Println(len(s1),cap(s1))
  fmt.Println(s1[0],s1[1],s1[2])
  s1=append(s1,10)
  fmt.Println(s1[0],s1[1],s1[2],s1[3])

}

//动态增加的cap

func TestSliceGrowing(t *testing.T)  {
  var s2 =[]int{}
  for i:=0;i<11;i++{
    s2=append(s2,i)
    fmt.Println(len(s2),cap(s2))
  }
}

//slice 是共享空间

func TestSliceShare(t *testing.T)  {
  year:=[]string{"1月","2月","3月","4月","5月","6月","7月","8月","9月","10月","11月","12月"};
  d1:=year[3:6];
  d2:=year[5:8];
  fmt.Println(d1,d2)
  d2[0]="不知道几月"
  fmt.Println(d1,year)

}
