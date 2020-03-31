package main

import "fmt"

func removeDuplicates(input []string) []string {
	i := 0
	for _, value := range input {
		if i != (len(input)) {
			if value != input[i+1] {
				input[i] = value
				i++
			}
		}
	}
	return input[:i]
}

func main() {
	var stuff []string
	stuff = append(stuff, "apple")
	stuff = append(stuff, "apple")
	stuff = append(stuff, "banana")
	stuff = append(stuff, "kiwi")
	fmt.Println(stuff)
	stuff = removeDuplicates(stuff)
	fmt.Println(stuff)
}
