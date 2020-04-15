package array1

import (
"fmt"
)

type Student struct {
	id uint
	name string
	age int
}

func newStudent(id uint,name string,age int) *Student {
	return &Student{id:id,name:name,age:age}
}

func (s Student) getName() string  {
	return s.name
}

func (s *Student) setName(newName string)  {
	s.name=newName;
}


func main()  {
	student:=newStudent(12,"wuqinqiang",20)
	student.setName("curry")
	fmt.Println(student.name);
}

