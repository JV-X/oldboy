package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

//	type Addr struct {
//		Province string
//		City     string
//		Zhen     string
//	}
//
//	type Customer struct {
//		Name string `json:"name"`
//		Age  int    `json:"age"`
//		Addr
//	}
type Addr struct {
	Province string
	City     string
	Zhen     string
}
type Customer struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr
}

func main() {
	//jsonBytes, err := ioutil.ReadFile("/Users/jianwei/GolandProjects/oldboy/day04/customers.json")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(jsonBytes))
	//// var customers []map[string]interface{}
	//var customers []Customer
	//json.Unmarshal(jsonBytes, &customers)
	//fmt.Println(customers)
	//fmt.Println(reflect.TypeOf(customers))
	////fmt.Println(customers[1].Age)
	//
	//for i, customerStruct := range customers {
	//	//fmt.Println(i, customerStruct)
	//	fmt.Printf("序号%d 姓名：%s 年龄：%d 省份:%s 城市:%s\n", i, customerStruct.Name, customerStruct.Age, customerStruct.Province, customerStruct.City)
	//}
	jsonBytes, err := ioutil.ReadFile("/Users/jianwei/GolandProjects/oldboy/day04/customers.json")
	if err != nil {
		fmt.Println(err)
	}
	var customers []Customer
	json.Unmarshal(jsonBytes, &customers)
	fmt.Println(customers)
	fmt.Println(reflect.TypeOf(customers))
	fmt.Println(customers[1].Age)
	for i, customerStruct := range customers {
		fmt.Printf("序号%d 姓名：%s 年龄：%d 省份:%s 城市:%s\n", i, customerStruct.Name, customerStruct.Age, customerStruct.Province, customerStruct.City)
	}
}
