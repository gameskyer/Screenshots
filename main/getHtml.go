package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.runoob.com/sqlite/sqlite-tutorial.html")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "findlinks1:%v\n", err)
		os.Exit(1)
	}
	//b, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(b))
	doc, err := html.Parse(resp.Body)
	//defer resp.Body.Close()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "findlinks1:%v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	fmt.Println(n.Data)
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
