package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	baseURL   = "https://xkcd.com/"
	numComics = 3
	extension = "/info.0.json"
)

//Returned from xkcd website json vals
type Comic struct {
	URL        string
	Title      string `json:"title"`
	Year       string `json:"year"`
	Num        int    `json:"num"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
}

//Collection of comics
type ComicList struct {
	Comics []*Comic
}

//num is page number
func downloadComic(num int) (*Comic, error) {
	url := baseURL + strconv.Itoa(num) + extension
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("ERROR: %s", resp.Status)
	}
	var result Comic
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func downloadComics() (*ComicList, error) {
	var result ComicList
	for i := 1; i < numComics; i++ {
		comic, err := downloadComic(i)
		if err != nil {
			return nil, err
		}
		result.Comics = append(result.Comics, comic)
	}
	return &result, nil
}

func downloadImage(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func main() {
	comics, _ := downloadComics()
	for _, c := range comics.Comics {
		imgURL := c.URL
		downloadImage(imgURL)
	}
}
