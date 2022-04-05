package main

import (
	"fmt"
	"reflect"
)

//Declare inspection function using interface
func inspector(x interface{}) {
	reflection := reflect.TypeOf(x)
	refValue := reflect.ValueOf(x)

	fmt.Println("")
	fmt.Println(x)
	fmt.Println(reflection.Name())
	fmt.Println(reflection.Kind())
	fmt.Println(reflection.NumField())

	for i := 0; i < reflection.NumField(); i++ {
		fmt.Println("")
		fmt.Println(refValue.Field(i))
		fmt.Println(reflection.Field(i).Name)
		fmt.Println(reflection.Field(i).Type)
	}
}

func main() {
	//declare object
	type customer struct {
		fName string
		lName string
		userID int
		invoiceTotal float64
	}

	customer1 := customer{"John", "Wick", 123123123, 10000.0}

	//trigger inspection
	inspector(customer1)
}
