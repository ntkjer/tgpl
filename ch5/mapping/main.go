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
	for k, v := range visit(make(map[string]int), doc) {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func visit(freq map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return freq
	}
	if n.Type == html.ElementNode {
		freq[string(n.Data)]++
	}

	//Traverse next node
	//for c := n.FirstChild; c != nil; c = c.NextSibling {
	//	freq = visit(freq, c)
	//}
	//The following is the same but is more concise
	visit(freq, n.FirstChild)
	visit(freq, n.NextSibling)

	return freq
}
