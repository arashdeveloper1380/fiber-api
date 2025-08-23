package pkg

import (
	"fmt"
	"reflect"
)

func CheckVariableWithReflect() {
	x := 40
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println("Type:", t)        // int
	fmt.Println("Kind:", t.Kind()) // int
	fmt.Println("Value:", v)       // 42

}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func CheckStructFieldReflect() {
	p := Person{"arash", 24}
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	fmt.Println("Struct Name is :", t.Name())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fmt.Printf("Field: %s, Type: %s, Value: %v\n", field.Name, field.Type, value)
	}
}

func add(a, b int) int {
	return a + b
}

func CallFuncDynamicReflect() {
	fn := reflect.ValueOf(add)
	args := []reflect.Value{reflect.ValueOf(5), reflect.ValueOf(2)}
	result := fn.Call(args)
	fmt.Println("Result:", result[0].Int())
}

func GetStructPropertyTags() {
	person := Person{Name: "ali", Age: 25}
	t := reflect.TypeOf(person)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field : %s, Json Tag is : %s\n", field.Name, field.Tag.Get("json"))
	}
}
