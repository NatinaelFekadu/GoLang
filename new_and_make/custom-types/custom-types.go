package customtypes

import "fmt"

type person struct {
	name string
	age  int
}

type customeNumber int

func (number customeNumber) pow(power int) customeNumber{
	result := number
	for i := 0; i < power; i++ {
		result *= number
	}
	return result
}

type personData map[string]person

func Main2() {
	var people personData = personData{"p1": {"Nathy", 17}}
	fmt.Println(people)

	var dummyNumber customeNumber = 4
	powerdNumber := dummyNumber.pow(3)
	fmt.Println(powerdNumber)
	
}