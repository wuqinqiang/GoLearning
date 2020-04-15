package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("发送方发送数据：%d\n", i)
			ch <- i
		}
		fmt.Printf("发送方: 关闭通道\n")
		close(ch)
	}()

	for {
		num, ok := <-ch
		if !ok {
			fmt.Printf("接收方: 通道已关闭\n")
			break
		}
		fmt.Printf("接收方接收到的数据是:%d\n", num)
	}
}
