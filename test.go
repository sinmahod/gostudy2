package main

import "fmt"

func Test1() {
	values := []int{9, 3, 4, 1, 2, 34, 7, 212}
	Sort(values)
	fmt.Println(values)
}

//测试结构体，实现排序功能，从大到小
type tree struct {
	value int
	max   *tree
	min   *tree
}

func Sort(values []int) {
	var t *tree
	for _, v := range values {
		t = add(t, v)
	}
	getValues(t, values[:0])
}

func getValues(t *tree, values []int) []int {
	/*
		      1
		     /  \
		   3    2
		  /  \
		9    4    7
		  \       /
		     34
		     \
		     212
	*/
	if t != nil {
		fmt.Println(t.value)
		values = getValues(t.max, values)
		values = append(values, t.value)
		values = getValues(t.min, values)
	}

	return values
}

func add(t *tree, v int) *tree {
	if t == nil {
		t = new(tree)
		t.value = v
		return t
	}
	if t.value > v {
		t.min = add(t.min, v)
	} else {
		t.max = add(t.max, v)
	}
	return t
}

type Ponit struct {
	X, Y int
}

type Circle struct {
	Ponit
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func Test2() {
	var w Wheel
	w = Wheel{Circle: Circle{Ponit: Ponit{X: 9, Y: 8}, Radius: 7}, Spokes: 6}
	fmt.Printf("%#v\n", w)
}

func main() {
	Test2()
}
