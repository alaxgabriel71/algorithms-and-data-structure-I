package main

import (
	"fmt"
	"local-packages/list"
)

func main() {
	array := list.ArrayList{}

	array.Init()
	// array.Add(0)
	// array.Add(1)
	// array.Add(2)
	// array.Add(3)
	// array.AddOnIndex(99,1)
	// array.AddOnIndex(88,5)
	// array.AddOnIndex(77,7)
	// array.RemoveOnIndex(1)
	// array.RemoveOnIndex(5)
	// array.RemoveOnIndex(4)
	// array.RemoveOnIndex(0)
	// array.RemoveOnIndex(-1)
	fmt.Println(array)
	fmt.Println(array.Size())
}