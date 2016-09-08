package test3

import (
	"fmt"
	"os"
)

func Test1() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		fmt.Printf("%c\n", x)
	}
}

//得到当前工作空间目录
func Test2() {
	cwd, err := os.Getwd()
	fmt.Println(cwd, err)
}
