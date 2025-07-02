package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("App_name"))
	var a1 int = 10
	fmt.Print(a1)
	var (
		a2 int
		a3 string
		a4 float32
	)
	a2, a3, a4 = 20, "Hello", 3.4
	fmt.Println(a2, a3, a4)
}
