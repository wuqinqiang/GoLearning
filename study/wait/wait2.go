package main

import (
	"fmt"
	"sync"
)

func add(a, b int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()

	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
}

func main() {
	total := 10
	step := 2
	fmt.Println("子协程启动")
	var wg sync.WaitGroup
	for i := 0; i < total; i = i + step {
		wg.Add(step)
		for j := 0; j < step; j++ {
			go add(i+j,1,wg.Done)
		}
		wg.Wait()
	}

	fmt.Println("所有子协程执行完毕")

}
