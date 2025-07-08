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
	//const pi = 3.14
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
	fmt.Println("arr1=", arr1[2])
	// thay doi gia tri mang
	arr1[0] = 10
	fmt.Println("arr1=", arr1[0])
	//khoi tao mang
	arr2 := [3]int{}
	arr3 := [3]float32{1, 2, 3}
	fmt.Print("arr2=", arr2)
	fmt.Print("arr3=", arr3)
	// chi tao phan tu cu the
	arr4 := [3]int{1: 10, 2: 30}
	fmt.Print("arr3=", arr4)
	// do dai phan tu cua mang
	fmt.Println(len(arr4))
	// go slices
	slices1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slices=%v \n", slices1)
	fmt.Println("length of slices1 =", len(slices1))
	fmt.Println("cap slices1=", cap(slices1))
	// tao lat cat trong mang
	arr65 := [6]int{101, 112, 122, 132, 142, 152}
	myslicess := arr65[2:4]
	fmt.Printf("myslice = %v\n", myslicess)
	fmt.Printf("length = %d\n", len(myslicess))
	fmt.Printf("capacity = %d\n", cap(myslicess))
	// tao slices bang make
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// with omitted capacity
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))
}
