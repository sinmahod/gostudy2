package main

import (
	"fmt"
	"gostudy2/people"
)

//练习方法使用
func Test1() {
	p := people.People{
		Age:  18,
		Name: "Lili",
	}
	fmt.Println(p.GetBirthYear())
	p.AddAge()
	fmt.Println(p.Age)
}

func Test2() {
	mp := make(map[string]int)
	mp["ccc"] = 123
	mp["ddd"] = 234
	fmt.Printf("%p\n", mp)
	mp2 := *mp
	mp["aaa"] = 987
	fmt.Println(mp2)
	fmt.Println(mp)
	fmt.Printf("%p", mp2)
}

func main() {
	Test2()
}
