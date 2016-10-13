package main

import (
	"bytes"
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

//测试bit数组

type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte('}')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func Test5() {
	var is IntSet
	fmt.Println(is.Has(2))
	is.Add(2)
	is.Add(3)
	is.Add(422)
	fmt.Println(is.Has(2))
	fmt.Println(is.String())

	fmt.Println(32 << (^uint(0) >> 63))
	fmt.Println(32 << 1)
}

func main() {
	Test5()
}
