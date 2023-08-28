package main

import "fmt"

func main() {
	arr := [4]int{10, 20, 30, 40}
	s1 := arr[0:2] // [10, 20]
	s2 := s1       //  // [10, 20]
	s3 := append(append(append(s1, 1), 2), 3)
	s1[0] = 1000
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(arr)
}
