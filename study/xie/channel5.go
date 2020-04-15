package main

import (
	"fmt"
	"time"
)

func main() {
	chas := make(chan int, 1)
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

	select {
	case <-chas:
		fmt.Printf("接收到ch的数据\n")
	case <-timeout:
		fmt.Printf("超时一秒，退出了\n")
	//default:
	//	fmt.Printf("程序出错了")
	}
}
