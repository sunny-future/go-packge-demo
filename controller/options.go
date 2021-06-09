package controller // 声明 包名 controller

import (
	"fmt"
)

func WithAddr() string {
	addr := "127.0.0.1"
	fmt.Println("it's module controller file options func WithAddr !")
	return addr
}