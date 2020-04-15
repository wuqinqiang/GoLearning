package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	chs:=[6]chan int{
		make(chan int,1),
		make(chan int,1),
		make(chan int,1),
		make(chan int,1),
		make(chan int,1),
		make(chan int,1),
	}

	index:=rand.Intn(6)
	fmt.Printf("随机数字数:%d\n",index)
	chs[index]<-index
	select {
	case <-chs[0]:
		fmt.Println("第一个分支被选中")
	case <-chs[1]:
		fmt.Println("第二个分支被选中")
	case <-chs[2]:
		fmt.Println("第三个分支被选中")
	default:
		fmt.Println("没有分支被选中")
	}

}
