package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int = 0

func add(a, b int, lock *sync.RWMutex) {

	c := a + b
	lock.Lock()
	counter++
	lock.Unlock()
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
}

func main() {
	start := time.Now()
	lock := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		go add(1, i, lock)
	}

	for {
		lock.RLock()
		res := counter
		lock.RUnlock()
		runtime.Gosched()
		if res >= 10 {
			break
		}
	}

	second := time.Now().Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", second)

}
