package main

import (
	"fmt"

	singleton "github.com/iamprahladsinghnegi/design-patterns/creational/singleton/singleton"
)

func main() {

	for i := 0; i < 30; i++ {
		go singleton.GetInstance()
	}

	fmt.Scanln()
}
