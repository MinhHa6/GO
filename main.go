package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("App_name"))
	var a1 int = 10
	fmt.Print(a1)
}
