package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"study/json/common"
)

type ServiceHandle struct{}

func (ServiceHandle *ServiceHandle) GetName(id int, item *common.Item) error {
	log.Printf("receive GetName call,id:%d", id)
	item.Id = id
	item.Name = "curry"
	return nil
}

func (ServiceHandle *ServiceHandle) SaveName(item common.Item, resp *common.Response) error {
	log.Printf("receive SaveName call,Item:%v", item)
	resp.Ok = true
	resp.Id = item.Id
	resp.Message = "存储成功"
	return nil
}

func main() {
	server := rpc.NewServer()

	//监听端口8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("监听端口失败：%v", err)
	}
	defer listener.Close()

	log.Println("Start listen on port localhost:8080")
	// 初始化服务处理器
	serverHandle := &ServiceHandle{}
	err = server.Register(serverHandle)
	if err != nil {
		log.Fatalf("注册服务处理器失败:%v", err)
	}

	//等待客户端连接
	for {
		conn, error := listener.Accept()
		if error != nil {
			log.Fatalf("接收客户端连接失败:%v", error)
		}
		//自定义RPC编码器，新建一个jsonrpc 编码器
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
