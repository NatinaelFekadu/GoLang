package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func (prod *Product) store(){
	file, _ := os.Create(prod.id + ".txt")
	content := fmt.Sprintf(
		"ID:%v\nTitle:%v\nDescription:%v\nPrice: USD %.2f\n",
		prod.id,
		prod.title,
		prod.description,
		prod.price,
	)
	file.WriteString(content)

	file.Close()
}

// reader := 

func main() {
	createdProduct := getProduct()

	createdProduct.outputData()

	createdProduct.store()
}

func NewProduct(id string, title string, description string, price float64) *Product{
	newProduct := Product{
		id,
		title,
		description,
		price,
	}
	return &newProduct
}

func (product *Product) outputData(){
	fmt.Printf("ID:%v\n",  product.id)
	fmt.Printf("Title:%v\n",  product.title)
	fmt.Printf("Description:%v\n",  product.description)
	fmt.Printf("Price: USD %.2f\n",  product.price)
}

func getProduct() *Product{
	reader := bufio.NewReader(os.Stdin)
	idInput := readUserInput(reader,"Product ID: ")
	titleInput := readUserInput(reader,"Product Title: ")
	descriptionInput := readUserInput(reader,"Product Description: ")
	priceInput := readUserInput(reader,"Product Price: ")

	priceValue,_ := strconv.ParseFloat(priceInput,64)

	product := NewProduct(idInput,titleInput,descriptionInput,priceValue)
	
	return product
}

func readUserInput(reader *bufio.Reader,promtpText string) string{
	fmt.Print(promtpText)
	userInput,_ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	return userInput
}
