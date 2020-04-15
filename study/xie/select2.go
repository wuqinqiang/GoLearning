package main

import (
	"fmt"
	"math/rand"
)

func main() {
	chs := [3]chan int{
		make(chan int, 3),
		make(chan int, 3),
		make(chan int, 3),
	}

	res:=rand.Int()

	fmt.Print(res)

	index1 := rand.Intn(100)
	fmt.Printf("随机数字1:%d\n", index1)

	chs[0] <- rand.Int()

	index2 := rand.Intn(100)
	fmt.Printf("随机数字2:%d\n", index2)

	chs[1] <- rand.Int()

	index3 := rand.Intn(100)
	fmt.Printf("随机数字3:%d\n", index3)
	chs[2] <- rand.Int()
	for i := 0; i < 3; i++ {
		select {
		case num, ok := <-chs[0]:
			if !ok {
				break;
			}
			fmt.Println("第一个条件分支被选中:chs[0]=>", num)
		case num, ok := <-chs[1]:
			if !ok {
				break;
			}
			fmt.Println("第二个分支被选中:chs[1]=>", num)

		case num, ok := <-chs[2]:
			if !ok {
				break;
			}
			fmt.Println("第三个分支被选中:chs[2]=>", num)

		default:
			fmt.Println("没有被选中")


		}
	}

}
