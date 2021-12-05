package main

import "fmt"

func main() {
	return abc() efg() cde()
}
func abc() {
	s := [...]string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "apple", "banana" and "coconut" needed
	 s = s[0:2]
	fmt.Print(s)
	// [apple banana coconut]
}

func efg() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "elderberries", "figs" and "grapes" needed
	 s  = s[4:6]
	fmt.Print(s)
	// * [elderberries figs grapes]
}

func cde() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "coconut", "durian" and "elderberries" needed
	 s = s[2:4]
	fmt.Print(s)
	// * [coconut durian elderberries]
}