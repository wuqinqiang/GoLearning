package common

type Item struct {
	Id   int `json:"id"` 
	Name string `json:"name"`
}

type Response struct {
	Ok bool `json:"ok"`
	Id int `json:"id"`
	Message string `json:"msg"`
}




