package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html"
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

//测试模板
type Html struct {
	TotalCount int
	Items      []Issc
}

type Issc struct {
	Number    int
	User      *User
	Title     string
	CreatedAt time.Time
	Test      template.HTML
}

type User struct {
	Login bool
}

const templ = `<html><title>测试</title><body>{{.TotalCount}} issues:
	{{range .Items}}----------------------------------------
	Number: {{.Number}}
	<h2>User:   {{.User.Login}}</h2>
	<hr/>Title:  {{.Title | printf "%.64s"}}
	<h1>Age:    {{.CreatedAt | daysAgo}} days</h1>
	<h1>Test:    {{.Test}} </h1>
	{{end}}</body></html>`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func Test5(w http.ResponseWriter, r *http.Request) {
	text := template.Must(template.New("test").
		Funcs(template.FuncMap{"daysAgo": daysAgo}). //注册函数
		Parse(templ))                                //放入模板

	t1, _ := time.Parse("2016-08-08 20:47:19", "2015-06-06 20:47:19")
	t2, _ := time.Parse("2016-06-06 20:47:19", "2015-06-06 20:47:19")
	isr := Html{TotalCount: 10, Items: []Issc{
		{
			Number: 5,
			User: &User{
				Login: false,
			},
			Title:     "<b>test1</b>",
			CreatedAt: t1,
			Test:      "<b>test1<b>",
		}, {
			Number: 6,
			User: &User{
				Login: true,
			},
			Title:     "test2",
			CreatedAt: t2,
			Test:      "<b>test2<b>",
		},
	},
	}
	text.Execute(w, isr)

	//method main
	//http.HandleFunc("/", Test5)
	//http.ListenAndServe("localhost:8080", nil)
}

//遍历所有A标签打印所有链接
func Test6() {
	if file, err := os.Open("index.html"); err == nil {
		if doc, err := html.Parse(file); err == nil {
			for i, s := range visit(nil, doc) {
				fmt.Println(i, s)
			}
		}
	}
	//测试URL
	// if resp, err := http.Get("http://www.baidu.com/"); err == nil {
	// 	if doc, err := html.Parse(resp.Body); err == nil {
	// 		//parseHtml(doc)
	// 		for i, s := range visit(nil, doc) {
	// 			fmt.Println(i, s)
	// 		}
	// 	}
	// 	defer resp.Body.Close()
	// }
}

func visit(links []string, n *html.Node) []string {

	if n.Type == html.ElementNode {
		fmt.Println(n.Data)
		// for _, a := range n.Attr {
		// 	if a.Key == "href" {
		// 		links = append(links, a.Val)
		// 	}
		// }
	}
	//c = c.NextSibling

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}

func parseHtml(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.Data == `a` {
			for _, a := range n.Attr {
				if a.Key == `href` && strings.HasPrefix(a.Val, "http:") {
					fmt.Println("+", a.Val, "+")
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseHtml(c)
	}
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

//练习5.2
func Test7() {
	if file, err := os.Open("index.html"); err == nil {
		if doc, err := html.Parse(file); err == nil {
			mp := make(map[string]int)
			total(mp, doc)
			//引用类型传递的时候都是传递的地址，所以这里不需要返回
			for k, v := range mp {
				fmt.Println(k, v)
			}
		}
	}
}

func total(mp map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		mp[n.Data]++
	}

	if n.FirstChild != nil {
		total(mp, n.FirstChild)
	}
	if n.NextSibling != nil {
		total(mp, n.NextSibling)
	}
}

//练习5.3
func Test8() {
	if file, err := os.Open("index.html"); err == nil {
		if doc, err := html.Parse(file); err == nil {
			visit2(nil, doc)
		}
	}
}

func visit2(links []string, n *html.Node) []string {

	if n.Type == html.ElementNode && n.Data == `a` {
		fmt.Println(n.FirstChild.Data) //获取<a>标签</a>之间的内容
	}
	//c = c.NextSibling

	if n.FirstChild != nil {
		links = visit2(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit2(links, n.NextSibling)
	}

	return links
}

func main() {
	Test8()
}
