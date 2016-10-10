package people

import "time"

type People struct {
	Age  int
	Name string
}

//对象型，不改变值，一般情况下如果有指针型所有的方法都必须为指针型（规范）
func (p People) GetBirthYear() int {
	return time.Now().Year() - p.Age
}

//指针型，可以改变内存对象的值
func (p *People) AddAge() {
	p.Age++
}
