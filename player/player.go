package player

import (
	"fmt"
)

type Opr interface {
	GetInput() (OPT, bool)
}

// 操作人
type Operator struct {
	name string
}
type OPT int

// 定义操作的枚举
const (
	UP OPT = iota
	DOWN
	LEFT
	RIGHT
	NOTHING
)

func getName() string {
	return "CanYang"
}
func NewOpr() *Operator {
	getName()
	opr := &Operator{
		name: getName(),
	}
	return opr
}

// 获取操作的函数
func (opr *Operator) GetInput() (opt OPT, flag bool) {
	var input string
	fmt.Scanln(&input)
	//获取用户的操作
	if input == "w" {
		fmt.Println("上")
		return UP, true
	}
	if input == "a" {
		fmt.Println("左")
		return LEFT, true
	}
	if input == "d" {
		fmt.Println("右")
		return RIGHT, true
	}
	if input == "s" {
		fmt.Println("下")
		return DOWN, true
	}
	return NOTHING, false
}
