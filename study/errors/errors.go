package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("保底运行")
	}()
	p := 1
	q := 0
	if q == 0 {
		panic("除数不能为0")
	}

	c:=p/q;

	fmt.Println("得到%f",c);

}
