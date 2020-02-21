package array_test
import(
  "testing"
  "fmt"
)

//数组
func TestArrayInit(t *testing.T)  {
  var arr [3]int

  arr[0]=1;
  fmt.Println(arr[0],arr[1])
   arr1:=[5]int{1,2,3,4,5}
   arr2:=[...]int{7,8,9}

   fmt.Println(arr1,arr2)

}

func TestArrayTravel(t *testing.T)  {
  arr1:=[5]int{10,20,30,40,50}
  //数组遍历
  // for index,e:=range arr1{
  //   fmt.Println(index,e)
  // }
  //不关心索引位置的话,用下划线
  for _,e:=range arr1{
    fmt.Println(e)
  }
}

func TestArrayMany(t *testing.T)  {
  c:=[...][2]int{{3,4},{5,6}}
  fmt.Println(c)
}

//截取数组
func TestArraySub(t *testing.T)  {
  arr:=[...]int{10,20,30,40,50}
  arr2:=arr[1:]
  arr3:=arr[2:4]

  arr4:=arr[:3]

  fmt.Println(arr2,arr3,arr4)

}
