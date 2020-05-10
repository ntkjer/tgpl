package main

import "fmt"

func main() {
	fmt.Println(weird())
}

func weird() (ret string) {
	defer func() {
		recover()
		ret = "hello world"
	}()
	panic("panic!")
}
