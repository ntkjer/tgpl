package countwordsandimages

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("Parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	words = 0
	images = 0
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		fmt.Println("type ", n.Type)
		if n.Data == "image" {
			images++
		}
	}
	return words, images
}
