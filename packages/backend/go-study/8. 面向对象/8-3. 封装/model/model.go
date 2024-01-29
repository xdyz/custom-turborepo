package model

import (
	"errors"
)

type pserson struct {
	Name string
	age  int
}

// func (p *pserson) SetAge(a int) (int, error) {
// 	if a < 0 {
// 		return 0, errors.New("age must be greater than 0")
// 	} else {
// 		p.age = a
// 	}

// 	return a, nil
// }

func (p *pserson) SetAge(a int) {
	if a < 0 {
		errors.New("age must be greater than 0")
	} else {
		p.age = a
	}
}

func (p pserson) GetAge() int {
	return p.age
}

// 写一个工厂函数，返回一个Person对象，相当于构造函数

func NewPerson(name string, age int) *pserson {
	return &pserson{
		Name: name,
		age:  age,
	}
}
