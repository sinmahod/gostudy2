package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

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

//测试JSON
type Data struct {
	//变量名大写才能导出
	Content string   `json:"content"` //类似别名
	Index   int      `json:"index"`
	Key     []string `json:"key,omitempty"` //omitempty选项表示为空值是不输出
}

func Test3() {
	var data = []Data{{Content: "这是一个json练习", Index: 999, Key: []string{"一", "二", "三"}}, {Content: "这是一个json练习", Index: 999}}

	//紧凑化json
	if json, err := json.Marshal(data); err != nil {
		log.Fatalf("%s\n", err)
	} else {
		fmt.Printf("%s\n", json)
	}
	//格式化json，前缀与每一行的缩进
	if json, err := json.MarshalIndent(data, "", "	"); err != nil {
		log.Fatalf("%s\n", err)
	} else {
		fmt.Printf("%s\n", json)
	}
}

type IssuesSearchResult struct {
	Total_count        int    `json:"total_count"`
	Incomplete_results bool   `json:"incomplete_results"`
	Sssdf              string `json:"sssdf"`
}

//测试json解码
func Test4() {

	geturl := "https://api.github.com/search/issues"
	str := os.Args[1:]                           //repo:golang/go is:open json decoder
	q := url.QueryEscape(strings.Join(str, " ")) //decode参数
	resp, err := http.Get(geturl + "?q=" + q)
	if err != nil { //https://api.github.com/search/issues?q=repo%3Agolang%2Fgo+is%3Aopen+json+decoder
		log.Fatalf("%s\n", err)
	}
	if resp.StatusCode != http.StatusOK { //200
		resp.Body.Close()
		fmt.Errorf("search query failed: %s", resp.Status)
		return
	}

	//打印内容
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		fmt.Errorf("%s", err)
		return
	} else {
		var v struct {
			Total_count int `json:"total_count"`
		}
		if err := json.Unmarshal(body, &v); err != nil {
			fmt.Errorf("%s", err)
		}
		fmt.Println(v)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		fmt.Errorf("search query failed: %s", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("%d\n", result.Total_count)
	fmt.Printf("%b\n", result.Incomplete_results)
	fmt.Printf("%s\n", result.Sssdf)
	fmt.Println(result)

	//for _, item := range result.Items {
	//	fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	//}
}

func main() {
	Test4()
}
