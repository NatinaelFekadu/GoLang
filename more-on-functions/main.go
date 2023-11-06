package main

import (
	"fmt"

	"github.com/none/more-on-functions/recursion"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	tripleFn := createTransformer(3)
	doubled := transformNumbers(&numbers, getTransformerFuction(&numbers))
	// tripled := transformNumbers(&numbers, func(number int)int{ return number * 3})
	tripled := transformNumbers(&numbers,tripleFn)
	fmt.Println(doubled)
	fmt.Println(tripled)
	fmt.Println(sumUp(3,4,3,2))
	fmt.Println(sumUp(numbers...))
	fmt.Println(recursion.Factorial(5))
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	tNumbers := []int{}
	for _, val := range *numbers{
		tNumbers = append(tNumbers, transform(val))
	}
	return tNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformerFuction(numbers *[]int)transformFn{
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func createTransformer(factor int) transformFn{
	return func(number int) int{
		return number * factor
	}
}

/* VARIADIC FUNCTION */
func sumUp(numbers ...int) int{
	sum := 0
	for _,val := range numbers{
		sum += val
	}
	return sum
}