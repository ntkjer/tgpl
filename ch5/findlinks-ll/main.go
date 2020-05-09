package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	//	for _, link := range visit(nil, doc) {
	//		fmt.Println(link)
	//	}
	visit(nil, doc)
}

func visit(links []string, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
				fmt.Println(links)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(links, c)
	}
}
