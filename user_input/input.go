package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/none/input/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMatrices() (new_value1 float64, new_value2 float64,err1 error,err2 error){
	new_value1,err1 = getUserInput(info.FirstValue)

	new_value2,err2 = getUserInput(info.SecondValue)

	return
}

func getUserInput(promptText string) (value float64,err error){
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	value, err = strconv.ParseFloat(userInput,64)


	return 

	/* FOR INTEGER INPUT */
		// value,_ = strconv.ParseInt(userInput,0,64)

}