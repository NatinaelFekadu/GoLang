package main

import (
	"fmt"

	customtypes "gith.com/none/advanced/custom-types"
)

func main() {
	numbers := make([]int, 2, 10)
	numbers[0] = 1
	numbers[1] = 2

	amap := make(map[string]string,5)
	amap["name"] = "Nathy"
	amap["occupation"] = "Student"

	

	for _,v := range amap {
		fmt.Println(v)
	}

	fmt.Println(cap(numbers))
	fmt.Println(numbers)
	numbers = append(numbers, 3, 4, 5)
	fmt.Println(numbers)

	/* new function */

	num := new(int)
	fmt.Println(num)
	fmt.Println(*num)

	customtypes.Main2()
}