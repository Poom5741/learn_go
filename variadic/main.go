package main

import "fmt"

func main() {
	variadicString("a")
}

func variadicString(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
