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
	fmt.Println(p.GetBirthYear(2))
	p.AddAge()
	fmt.Println(p.Age)
}

//测试调用方法是否会改变map值

type MP string

func (mp *MP) addEEE() {
	*mp = "asdasd"
}

func Test2() {
	var mp MP
	mp = "qweqwe"
	fmt.Println(mp)
	//ss(mp)
	mp.addEEE()
	fmt.Println(mp)

}

func ss(mp map[string]int) {
	mp["eee"] = 999
}

//上层结构集成子结构的方法（字结构必须匿名）
func Test3() {
	ps := people.ManyPeople{
		people.People{18, "lili"},
		"xixi",
	}
	fmt.Println(ps.GetBirthYear(23))
}

//测试方法值与接收器
func Test4() {
	ppg := people.People.GetBirthYear
	pp := people.People{28, "lili2"}
	fmt.Println(ppg(pp, 10))

}

func main() {
	Test4()
}
