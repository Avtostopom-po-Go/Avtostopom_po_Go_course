package main

import "fmt"

// Что выведет код и почему?
func main() {
	v := 5
	p := &v
	fmt.Println(*p)

	changePointer(p)
	fmt.Println(*p)
}

func changePointer(p *int) {
	v := 3
	p = &v
}
