package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string
	Website string
	Age     uint
	Male    bool
	Skills  []string
}

func main() {
	user := User{
		"wuqinqiang",
		"dddd",
		18,
		true,
		[]string{"PHP", "Go", "Java"},
	}
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("json编码失败：%v\n", err)
		return
	}

	var user2 User

	err= json.Unmarshal(u,&user2)
	if err != nil {
		fmt.Printf("解码失败：%v\n", err)
		return
	}

	fmt.Printf("编码的格式：%s\n", u)

	fmt.Printf("json 解码结果：%#v\n", user2)

}
