package main

import (
	"fmt"
)

type Animal struct {
	name string
}

func (a Animal) Call() string  {

	return "动物的叫声"
}

func (a Animal) eating() string  {
	return "吃东西";
}


func (a Animal) find() string  {
	return "hahah";
}

type Dog struct {
	Animal
}

func (dog Dog) find() string  {
	return "找东西"
}

func main()  {
	animal:=Animal{"狗"}
	dog:=Dog{animal}
	fmt.Println(dog.Call())
	fmt.Println(dog.Animal.find())
}
