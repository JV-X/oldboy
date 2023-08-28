package main

import (
	"fmt"
	"reflect"
)

func main() {

	// （1） 整型

	// b： 比特位  _  00 01
	//            __  00  01  10 11
	//            ___ 000 001 010 011  100 101 110 111
	// B： Byte 字节  1B = 8b _ _ _ _ _ _ __  2的 八次方  256个数字

	/*

		var age uint8 //  uint8无负数 0-255
		age = 200
		fmt.Println(age)

		var x = 100
		fmt.Println(x)

		xType := reflect.TypeOf(x)
		fmt.Println(xType) // int*/

	// （2）浮点型
	var f1 float32
	f1 = 0.123456789
	fmt.Println(f1)

	var f2 float64
	f2 = 0.1234567890123456789
	fmt.Println(f2)

	var f3 = 3.14
	fmt.Println(reflect.TypeOf(f3)) // 默认float64

	var f4 = 6.2e-2 // 即使是整数用科学技术表示也是浮点型
	fmt.Println(f4, reflect.TypeOf(f4))

	var f5 = 6.2e+1
	fmt.Println(f5, reflect.TypeOf(f5))

}
