package main

import (
	"fmt"
	"github.com/adambarreiro/list"
)

func main()  {
	l := list.New()

	l.Prepend("Perico palotes")
	fmt.Println(l.Head())
	fmt.Println(l.Size())

	value, _ := l.Get(0)
	fmt.Println(value)

	_, err := l.Get(1)
	fmt.Println(err)
}
