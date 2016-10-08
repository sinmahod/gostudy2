package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"sort"
	"time"
)

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val) //以URL的形式返回
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func Test1() {
	var rmdirs []func()
	str := []string{"123", "234", "345", "abc", "qwe"}
	for _, d := range str {
		d := d //至关重要，上面的range是依次给d赋值，但是d始终是同一个内存地址，这句的目的就是重新生成一个内存地址存放d
		fmt.Println(1, d)
		rmdirs = append(rmdirs, func() { //这里的d引用的是内存地址，如果没有上面的d := d那么这里引用的内存地址就是同一个内存地址，循环执行完毕的时候内存地址对应的值是qwe，并且每次被追加1
			d += "1"
			fmt.Println(d)
		})
	}
	for _, rmdir := range rmdirs {
		rmdir()
	}
}

//练习可变参数
func Test2() {
	vallist(1, 2, 3, 4, 5)
	cs := []int{1, 2, 3, 4, 5}
	vallist(cs...)
	fmt.Printf("%T", vallist)
}

func vallist(cs ...int) {
	fmt.Println(len(cs))
	fmt.Println(cap(cs))
	fmt.Printf("%T\n", cs)
	for _, c := range cs {
		fmt.Println(c)
	}
}

//练习5.15
func Test3() {
	cs := []int{234, 45, 12, 435, 231}
	if err := min(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if err := max(cs...); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(cs)
}

func min(cs ...int) error {
	if len(cs) == 0 {
		return errors.New(fmt.Sprintf("Error：cs length is 0"))
	}
	sort.Ints(cs)
	fmt.Println(cs[0])
	return nil
}

func max(cs ...int) error {
	if len(cs) == 0 {
		return errors.New(fmt.Sprintf("Error：cs length is 0"))
	}
	sort.Ints(cs)
	fmt.Println(cs[len(cs)-1])
	return nil
}

//练习5.16
func Test4() {
	fmt.Printf("%s\n", stringsort(",", "aaa", "bbb", "ccc"))
}

func stringsort(join string, strs ...string) string {
	var s string
	i := len(strs)
	idx := 1
	if i > 0 {
		s += strs[idx]
	}
	var sfunc func()
	sfunc = func() {
		s += join + strs[idx]
		idx++
		if idx != i {
			sfunc()
		}
	}
	sfunc()
	return s
}

//练习defer语句，结论：defer最后执行强行中断不执行。
func Test5() {
	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	defer fmt.Println(22)
	//os.Exit(0)
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the
	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func double(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	}()
	return x + x
}

//执行顺序：result := double(x) ; defer ; return result  先得到result在defer计算，最后才返回，所以得到的值为 x + x + x
func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

//测试bath.Base函数
func Test6() {
	resp, err := http.Get("http://60.205.164.3/test/qwe")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	str := path.Base(resp.Request.URL.Path)
	fmt.Println(str)
	f, err := os.Create(str)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

//测试panic和recover , 练习5.19
func Test7() {
	// if err := getErr(); err != nil {
	// 	fmt.Printf("%s\n", err) //Error!!!,asdsd
	// }
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("%s,%s", "Error!!!", p)
		}
	}()
	getErr()
}

func getErr() {
	panic("asdsd")
}

func main() {
	Test7()
}
