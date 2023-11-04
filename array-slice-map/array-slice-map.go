package main

import "fmt"

// func main() {
// 	// var values [6]float64 = [6]float64{3,4,5,6,7}
// 	// values := [6]float64{3, 4, 5, 6, 7, 8}
// 	// valuesSlice := values[:4]

// 	// fmt.Println(values)
// 	// fmt.Println(len(valuesSlice),cap(valuesSlice))
	
// 	// valuesSlice = valuesSlice[:6]
// 	// fmt.Println(valuesSlice)

// 	/* DYNAMIC ARRAYS */
// 	values := []int{3,4,5}
// 	values = append(values, 6)
// 	otherValues := []int{7,8,9}
// 	values = append(values, otherValues...)
// 	fmt.Println(values)

// }

func main() {
	websites := map[string]string{
		"Google":             "https://google.com",
		"Amazon Web Service": "https://aws.com",
	}

	websites["LinkedIn"] = "https://linkedin.com"
	delete(websites,"Google")
	fmt.Println(websites)
}