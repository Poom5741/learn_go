package main

import (
	"fmt"

	"github.com/Poom5741/pack/book"
)

func main() {
	b := book.New()

	fmt.Printf("%T\n", b)
}
