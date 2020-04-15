package main

import (
	"fmt"
)

func add(a,b int)  {
	var c=a+b
	fmt.Println(c)
}

func main()  {
	for i:=1; i<10;i++  {
		go add(1,i)

	}
	//time.Sleep(1e9)
}
