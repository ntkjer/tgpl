package main

import (
	"bufio"
	"log"
	"os"
)

//Reports the freqeuncy of various words in an input text file
func wordFreq(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//Take freq count k[word] : v[count]
	counts := make(map[string]int, len(lines))
	for _, line := range lines {
		counts[line]++
	}
	log.Println(counts)
}

func main() {
	if len(os.Args) > 2 {
		log.Fatalf("Please only enter two args")
	}
	path := os.Args[1]
	wordFreq(path)

}
