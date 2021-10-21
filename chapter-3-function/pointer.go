package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println("포인터 값 자체: ", myIntPointer)
	fmt.Println("포인터 주소에 있는 값", *myIntPointer)

	var myFloat float64
	var myFloatPointer *float64
	myFloatPointer = &myFloat
	fmt.Println(reflect.TypeOf(&myFloat))
	fmt.Println(myFloatPointer)

	var myBool bool
	var myBoolPointer *bool
	myBoolPointer = &myBool
	fmt.Println(reflect.TypeOf(&myBool))
	fmt.Println(myBoolPointer)
}
