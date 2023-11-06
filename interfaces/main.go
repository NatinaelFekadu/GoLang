package main

import (
	"fmt"
	"io/fs"
	"os"
)

type logger interface{
	log()
}

type logData struct {
	message  string
	fileName string
}

func (ldata *logData) log() {
	err := os.WriteFile(ldata.fileName,[]byte(ldata.message),fs.FileMode(0644))
	if err != nil {
		fmt.Println("Storing the log data failed!")
	}
}

type loggableString string

func (text loggableString) log(){
	fmt.Println(text)
}

func main(){
	userLog := &logData{"User logged in","user-log.txt"}
	createLog(userLog)
	outputValue(userLog)

	message := loggableString("[INFO] User created!")
	createLog(message)
	outputValue(message)

	firstNumber := 5
	secondNumber := 10.2
	numbers := []int{1,2,3}

	dfn := double(firstNumber)
	dsn := double(secondNumber)
	dns := double(numbers)
	ds := double("Test")

	fmt.Println(dfn)
	fmt.Println(dsn)
	fmt.Println(dns)
	fmt.Println(ds)
}

func createLog(data logger){
	data.log()
}

func outputValue(value interface{}){
	fmt.Println(value)
}

func double(value interface{}) interface{}{
	switch val := value.(type){
	case int:
		return val * 2
	case float64:
		return val * 2
	case []int:
		newNumbers := append(val,val...)
		return newNumbers
	default:
		return ""
	}
}