package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	base()

	arr := []int{1, 2, 3}
	printArr(arr)

	// underlying()
	// fmt.Println("-------------------")

	// var m myMap = make(map[int]string)
	// m.add(1, "a")
	// fmt.Println(m) //map[1:a]
	// fmt.Println("-------------------")

	// inherit()
	// fmt.Println("-------------------")

	// inheritByInterface()
	// fmt.Println("-------------------")

	// embedding()
	// fmt.Println("-------------------")
}

func base() {
	var x int = 100

	var a interface{} = x
	var b interface{} = &x

	fmt.Printf("unsafe.Sizeof of a(%d) b(%d) reflect.ValueOf a(%#v) b(%#v)  \n",
		unsafe.Sizeof(a), unsafe.Sizeof(b), reflect.ValueOf(a), reflect.ValueOf(b))

	y := interface{}(x).(int)
	z := interface{}(x)

	fmt.Printf("x(%#v) y(%#v) z(%#v) \n",
		reflect.ValueOf(x), reflect.ValueOf(y), reflect.ValueOf(z))
}

func printArr(arr interface{}) {
	a, ok := arr.([]int)
	if ok {
		for _, v := range a {
			fmt.Printf("v(%v)\n", v)
		}
	}

	slice := reflect.ValueOf(arr)

	if slice.Kind() != reflect.Slice {
		panic(fmt.Sprintf("%+v is not a slice", slice))
	}
	fmt.Printf("slice.Slice(0, 0).Interface():(%+v)\n", slice.Slice(0, 0).Interface())
}
