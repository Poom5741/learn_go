package main

func main() {
	n := 2
	double(&n)
	name := "Bob"
	appendGreeting(&name)
}

// n := 2
// double(&n)
// n == 4
func double(n *int) {
	*n = *n * 2
}

// name := "Bob"
// appendGreeting(&name)
// name == "Hi, Bob"
// func appendGreeting(s *string) {
// 	var name string = ("Hi , %s" * s)
// 	return name
// }
