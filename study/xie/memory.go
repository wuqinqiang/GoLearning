package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int=0

func add(a ,b int,lock *sync.Mutex)  {
	c:=a+b
	lock.Lock()
	counter++
	fmt.Printf("%d+%d=%d\n",a,b,c)
	lock.Unlock()
}

func main()  {
	start:=time.Now()
	lock:=&sync.Mutex{}
	for i:=0;i<=10 ;i++  {
		go add(1,i,lock)
	}

	for {
		lock.Lock()
		count:=counter
		lock.Unlock()
		runtime.Gosched()
		if count>=10 {
			break
		}

	}
	end:=time.Now()

	consume:=end.Sub(start).Seconds()
	fmt.Println("程序耗时(s)：",consume)


}