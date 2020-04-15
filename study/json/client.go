package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"
	"study/json/common"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", "localhost:8080", 30*time.Second) //30秒超时时间
	if err != nil {
		log.Fatalf("客户端连接失败:", err)
	}

	defer conn.Close()

	client := jsonrpc.NewClient(conn)
	var item common.Item

	client.Call("ServiceHandle.GetName", 1, &item)
	log.Printf("ServiceHandle.GetName 返回结果:%v\n", item)
	var resp common.Response

	item = common.Item{2, "curry"}
	client.Call("ServiceHandle.SaveName", item, &resp)
	log.Printf("ServiceHandler.SaveName 返回结果：%v\n", resp)

}
