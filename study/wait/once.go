package main

import (
	"fmt"
	"sync"
	"time"
)

func dosome(o *sync.Once)  {
	fmt.Println("开始")
	o.Do(func() {
		fmt.Println("do someing....")
	})
	fmt.Println("结束")
}

func main()  {
	o:=&sync.Once{}
	go dosome(o)
	go dosome(o)
	time.Sleep(time.Second*1)
}
