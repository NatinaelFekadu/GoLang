package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n,m := generateRandomNumbers()
	fmt.Println(n)
	fmt.Println(m)
}

func generateRandomNumbers() (r1 int,r2 int){
	r1 = rand.Intn(10)
	r2 = rand.Intn(10)
	return
}