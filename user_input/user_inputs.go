package main

import "fmt"

func main() {

	new_value1, new_value2, err1, err2 := getUserMatrices()

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid Input!")
		return
	}

	prod := calculateProduct(new_value1, new_value2)

	printProduct(prod)

}

func calculateProduct(val1 float64, val2 float64) float64 {
	return val1 * val2
}