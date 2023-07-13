package main

import "fmt"

type studenter interface {
	Create(name string) student
}

type student struct {
	name string
}

type studenterImpl struct{}

func (s studenterImpl) Create(name string) student {
	return student{"doan" + name}
}

type maleStudent struct {
	student
	studenter
	gender string
}

func main() {
	stu := maleStudent{
		student:   student{},
		studenter: studenterImpl{},
		gender:    "male",
	}
	fmt.Println(stu.Create("dat"))
}
