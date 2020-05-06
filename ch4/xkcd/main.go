package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	random    = false
	imagePath = "."
	filter    = "jpg"
	baseURL   = "https://xkcd.com/"
	numComics = 15
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

type Image struct {
	Path string
	File *os.File
}

type LocalImages struct {
	Size   int
	Images map[int]*Image
}

func checkLocalImages() *LocalImages {

	var collection LocalImages
	collection.Images = make(map[int]*Image)
	collection.Size = 0

	files, err := ioutil.ReadDir(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	for n, f := range files {
		filename := strings.SplitAfter(f.Name(), ".")
		fileExtension := filename[len(filename)-1]
		if fileExtension == filter {
			tmp, err := os.Open(f.Name())
			if err != nil {
				log.Fatal(err)
			}

			img := Image{Path: f.Name(), File: tmp}
			collection.Images[n] = &img
			collection.Size++
			//fmt.Println(img, img.Path)
		}
	}

	return &collection

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
	fmt.Println("downloading: ", url)
	splitURL := strings.SplitAfter(url, "/")
	filename := splitURL[len(splitURL)-1]

	file, err := os.Create(string(filename))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	localImages := checkLocalImages()
	comics, _ := downloadComics()
	for n, c := range comics.Comics {
		if n > localImages.Size {
			imgURL := c.Img
			downloadImage(imgURL)
		}
	}
}
