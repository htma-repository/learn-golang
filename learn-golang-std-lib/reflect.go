package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	FirstName string `firstname:"Tama" isHuman:"true" required:"true"` // Struct Tag
	LastName  string `firstname:"Tama" isHuman:"true"`
	Grade     int    `firstname:"Tama" isHuman:"true" required:"true"`
}

func (s *Student) GetPropertyInfo() {
	reflectValue := reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Pointer {
		reflectValue = reflectValue.Elem()
	}

	reflectType := reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("Name :", reflectType.Field(i).Name)
		fmt.Println("Type data :", reflectType.Field(i).Type)
		fmt.Println("Value :", reflectValue.Field(i).Interface())
		fmt.Println("Struct tag :", reflectType.Field(i).Tag.Get("firstname"))
		fmt.Println("Struct tag :", reflectType.Field(i).Tag.Get("isHuman"))
		fmt.Println("")
	}

}

func isValid(data any) (result bool) {
	result = true
	var typesData reflect.Type = reflect.TypeOf(data)
	for i := 0; i < typesData.NumField(); i++ {
		var requiredTag = typesData.Field(i).Tag.Get("required")
		if requiredTag == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			result = value != ""
			if result == false {
				return result
			}
		}
	}

	return result
}

func main() {

	student1 := &Student{
		FirstName: "Hutama",
		LastName:  "Trirahmanto",
		Grade:     10,
	}

	student1.GetPropertyInfo()

	student2 := Student{
		FirstName: "Hutama",
		LastName:  "",
		Grade:     10,
	}

	fmt.Println(isValid(student2))

}
