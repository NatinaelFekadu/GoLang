package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	value,err := calculateSum()
	if err != nil {
		fmt.Println("INVALID INPUT")
		return
	}
	fmt.Println(value)
}

func calculateSum() (total float64,error error){
	reader := bufio.NewReader(os.Stdin)
	userInput,err := reader.ReadString('\n')

	error = err

	userInput = strings.TrimSpace(userInput)
	inputNumbers := strings.Split(userInput,",")

	total = 0

	for _,value := range inputNumbers{
		parsed,err := strconv.ParseFloat(value,64)
		total += parsed
		if err != nil{
			error = err
			break
		}
	}

	return


}