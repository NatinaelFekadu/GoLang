package main

import "fmt"

func main() {
	val := 34
	valAdd := &val
	fmt.Println(valAdd)
	fmt.Println(*valAdd)

	*valAdd = 32
	fmt.Println(val)

	doubledValue := double(valAdd)
	fmt.Println(doubledValue)
	fmt.Println(val)

}

func double(number *int) int{
	doubledValue := *number * 2
	*number = 1000
	return doubledValue
}