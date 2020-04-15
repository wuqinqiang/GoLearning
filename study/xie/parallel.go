package main

import (
	"fmt"
	"runtime"
	"time"
)

func add(seq int, ch chan int) {
	defer close(ch)
	sum := 0
	for i := 1; i < 10000000; i++ {
		sum += i
	}
	fmt.Printf("子协程%d的运行结果是:%d\n", seq, sum)
	ch <- sum
}

func main() {
	start := time.Now()
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(1)
	chs := make([]chan int, cpus)
	for i := 0; i < len(chs); i++ {
		chs[i]=make(chan int,1)
		go add(i, chs[i])
	}
	sum := 0

	for _, ch := range chs {
		res := <-ch
		sum += res
	}
	second := time.Now().Sub(start).Seconds()
	fmt.Printf("执行时间(s)：%f\n", second)
	fmt.Printf("最终运行结果:%d",sum)
}
