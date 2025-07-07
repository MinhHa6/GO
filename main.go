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
	//khai bao ngan
	Ha := "Helloo"
	// khai bao hang so
	fmt.Println(Ha)
	const pi = 3.14
	// if else
	var age int
	fmt.Println("Nhap tuoi cua ban:")
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("Bạn đã đủ tuổi")
	} else {
		fmt.Println("Bạn chưa đủ tuổi")
	}
	// for
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}
	// for range
	n := []int{1, 2, 4, 5, 76, 8} // slice kiểu []int	for i,v:=range n
	{
		fmt.Print("i=", n)
	}
	// swich
	switch day := 3; day {
	case 1:
		fmt.Println("Thứ Hai")
	case 2:
		fmt.Println("Thứ Ba")
	default:
		fmt.Println("Một ngày nào đó")
	}
	// while do while
	//go
	//  khai bao mang
	var arr1 = [3]int{1, 2, 3}
	fmt.Print("arr1", arr1[2])
}
