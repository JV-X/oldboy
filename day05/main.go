package main

import "fmt"

func f1() int {
	i := 5
	defer func() {
		i++
	}()
	return i
}

func f2() *int {

	i := 5
	defer func() {
		i++
		fmt.Printf(":::%p\n", &i)
	}()
	fmt.Printf(":::%p\n", &i)
	return &i
}

func f3() (result int) {
	defer func() {
		result++
	}()
	return 6
}

func f4() (result int) {
	defer func() {
		result++
	}()
	return result // ret result变量的值
}

func f5() (r int) {
	t := 5
	defer func() {
		t = t + 1
	}()
	return t
}

func f6() (r int) {
	defer func() {
		r = r + 1
	}()
	return 5
}

func f7() (r int) {
	defer func(x int) {
		r = x + 1
	}(r)
	return 5
}

func main() {
	//println(*f2())
	//println(f3())
	//println(f4())
	println(f7())
}
