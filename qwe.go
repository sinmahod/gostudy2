// Copyright  2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	//"net/http"
	"os"
)

func main() {
	if file, err := os.Open("index.html"); err == nil {
		if doc, err := html.Parse(file); err == nil {
			for i, s := range visit(nil, doc) {
				fmt.Println(i, s)
			}
		}
	}
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

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode { //&& n.Data == `a` {
		fmt.Println(n.Data)
		// for _, a := range n.Attr {
		// 	if a.Key == "href" {
		// 		links = append(links, a.Val)
		// 	}
		// }
	}
	links = visit(links, n.NextSibling)
	links = visit(links, n.FirstChild)
	// links = visit(links, c)
	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	links = visit(links, c)
	// }
	return links
}

//!-
