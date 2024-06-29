package main

import "fmt" //นำเอา package เข้ามาทำงาน

func main() {
	fmt.Println("Hello World")
	fmt.Println("-- ^v^ --")

	// Manual Type Declaration
	// var <name> <type>
	var first_name string
	first_name = "Jitlada"

	var month = 12.0

	var last_name string = "Yotinta"

	// Type Inference
	// <nickname>:=<value>
	nickname := "Kawfang"
	age := 20

	fmt.Printf("%s %s %s\n", first_name, last_name, nickname)
	fmt.Println("Age = ", age, "Years")

	// Static Var
	// const <name> <type>
	const year_born int = 2003
	fmt.Println("Born in Year = ", year_born)

	// Show value and type
	fmt.Printf("nickname value = %v type = %T\n", nickname, nickname)
	fmt.Printf("age value = %v type = %T\n", age, age)
	fmt.Printf("age + age = %v\n", age+age)
	fmt.Printf("mothe / age = %v\n", month/float64(age))

	// Input
	// fmt.Scanf(string_format, address_list)

	var friend string
	fmt.Scanf("%s", &friend)
	fmt.Printf("friend = %s \n", friend)

	// If-else
	if age >= 10 {
		fmt.Println("Old")
	} else {
		fmt.Println("Young")
	}

	// No need break
	var testSwitch = 21
	switch testSwitch {
	case 20 : 
		fmt.Println("20")
	case 21 : 
		fmt.Println(21)
	default: 
		fmt.Println("> 21")
	}
}
