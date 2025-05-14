package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	sonuc := add(2, 3)
	fmt.Println("2 + 3 =", sonuc)
}
