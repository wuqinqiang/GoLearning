package main

import (
	"fmt"
	"log"
	"net/rpc"
	"study/rpc/utils"
)

func main() {
	var serverAddress = "localhost"
	client, err := rpc.DialHTTP("tcp", serverAddress+":8080") //同步调用
	if err != nil {
		log.Fatal("建立与服务端连接失败", err)
	}

	args := &utils.Args{10, 10}
	var replay int

	err = client.Call("MathService.Multiply", args, &replay)
	if err != nil {
		log.Fatal("调用远程方法 MathService.Multiply 失败:", err)
	}

	fmt.Printf("%d*%d=%d\n",args.A,args.B,replay)

	//异步调用
	divideCall:=client.Go("MathService.Divide",args,&replay,nil)
	for  {
		select {
		case <-divideCall.Done:
			fmt.Printf("%d/%d=%d\n", args.A, args.B, replay)
			return
		}
	}
}
