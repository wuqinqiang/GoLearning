package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

type DataBucket struct {
	buffer *bytes.Buffer
	mutex  *sync.RWMutex
	cond   *sync.Cond
}

func newDataBucket() *DataBucket {
	buf := make([]byte, 0)
	db := &DataBucket{
		buffer: bytes.NewBuffer(buf),
		mutex:  new(sync.RWMutex),
	}
	db.cond = sync.NewCond(db.mutex.RLocker())
	return db
}

//读取器
func (db *DataBucket) Read(i int) {
	db.mutex.RLock()
	defer db.mutex.RUnlock() //结束后释放锁
	var data []byte
	var d byte
	var err error
	for {
		//每次读取一个字节
		if d, err = db.buffer.ReadByte(); err != nil {
			if err == io.EOF { //缓冲区数据为空时执行
				if string(data) != "" { //data不为空，则打印它
					fmt.Printf("reader-%d:%s\n", i, data)
				}
			}
			db.cond.Wait() //缓冲区为空，那么等待
			data = data[:0]
			continue
		}
		data = append(data, d)
	}
}

func (db *DataBucket) Put(d []byte) (int, error) {
	db.mutex.Lock()         //打开写锁
	defer db.mutex.Unlock() //结束之后释放写锁
	//写入一个数据块
	n, err := db.buffer.Write(d)
	db.cond.Broadcast() //写入数据之后通过Broadcast 通知处于阻塞状态的读数据
	return n, err
}

func main() {
	db := newDataBucket()
	for i:=0;i<3 ;i++  {
		go db.Read(i) //开启读取器协程

	}
	for j:=0;j<10 ;j++  {
		go func(i int) {
			d := fmt.Sprintf("data-%d", i)
			db.Put([]byte(d))
		}(j)          //开启写入器协程
	}
	time.Sleep(100 * time.Microsecond)

}
