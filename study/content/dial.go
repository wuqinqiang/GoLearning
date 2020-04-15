package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

	}
	return result.Bytes(), nil
}

func main() {

	resp,err:=http.Get("https://xueyuanjun.com")
	if err !=nil{
		fmt.Printf("发送请求失败:%v",err)
		return
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout,resp.Body)


	os.Exit(0)

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	//从参数获取主机信息
	server := os.Args[1]

	//建立网络连接
	conn, err := net.Dial("tcp", server)
	checkErr(err)


	// 调用返回的连接对象提供的 Write 方法发送请求
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkErr(err)
	//调用返回的连接对象提供的Read方法读取所有响应数据
	result,err:=readFully(conn)
	checkErr(err)
	fmt.Println(string(result))
	os.Exit(0)
}
