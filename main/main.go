package main

import (
	"github.com/adambarreiro/list"
)

func main()  {
	l := list.New()
	l.Append(1).Append(2).Append(3).Append(4)
	l.Println()
}
