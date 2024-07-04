package main
// Package declaration

// Import packages
import "fmt" 

// Functions
func main() {
	// Printing in Go is similar to Python, but you need to write the package
	// Print is a function from fmt package
	fmt.Println("Hello World")
	fmt.Println("-- ^v^ --")

	// we can use ; to separated statement
	fmt.Print("Try to"); fmt.Printf("this %d", 030)

	// Comment with // or /* ... */ for many lines
	/* 	Comment 1 
		Comment 1 
		Comment 1 
		Comment 1 */

	// Variable Types
	/* 	int
		float32
		string
		bool */

	// Declare Variables

	// 'var' key : "Can be used inside and outside of functions" [https://www.w3schools.com/go/go_variables.php]
	// var <name> <type> = <value>
	var authorFName string = "Jitlada"
	var authorLName = "Yotinta"


	// ':=' key : "Can only be used inside functions" [https://www.w3schools.com/go/go_variables.php]
	// <name>:=<value>
	authorNName := "Kawfang"
	authorAge := 20

	fmt.Println("Author")
	fmt.Printf("Miss %s %s (%s)\n", authorFName, authorLName, authorNName)
	fmt.Println("Age = ", authorAge, "Years")

	// Multiple Variable Declaration
	var a, b, c int = 1, 2, 3
	d, f := 4, 5.0
	var (
		e int
		g int = 6
		h = 7.2
	  )

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("f = ", f)
	fmt.Println("e = ", e)
	fmt.Println("g = ", g)
	fmt.Println("h = ", h)

	// Constants Var
	// const <name> <type>
	const yearBorn int = 2003
	const (
		dayBorn int = 12
		monthBorn = 11
	)
	fmt.Println("Birthday:", dayBorn, monthBorn, yearBorn)

	// Show value and type
	fmt.Printf("nickname value = %v type = %T\n", authorNName, authorNName)
	fmt.Printf("age value = %v type = %T\n", authorAge, authorAge)
	fmt.Printf("age + age = %v\n", authorAge+authorAge)
	fmt.Printf("mothe / age = %v\n", dayBorn/authorAge)

	// Declare Array
	// var key
	/* 	"var array_name = [length]datatype{values} // here length is defined
		 var array_name = [...]datatype{values} // here length is inferred"
		[https://www.w3schools.com/go/go_arrays.php]
	*/

	var pets = [5] string {"Dog", "Cat", "Bird", "Tiger", "Lion"}
	var amountPet = [] int{10, 25, 26, 32, 10}
	var numPetCheck[5] int

	fmt.Println(pets, "Size =", len(pets))
	fmt.Println(amountPet)
	fmt.Println(numPetCheck)
	fmt.Printf("Type: %s Amount: %d\n", pets[4], amountPet[4])

	// := key
	/* 	"array_name := [length]datatype{values} // here length is defined
		 array_name := [...]datatype{values} // here length is inferred"
		[https://www.w3schools.com/go/go_arrays.php]
	*/

	days := [] string {"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	month := [12] string {"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	fmt.Println(days)
	fmt.Println(month)

	// Slices
	/* 	["https://www.w3schools.com/go/go_slices.php"]
		slice_name := []datatype{values}
		myslice := []int{}
	
		slice_name := make([]type, length, capacity)

		var myarray = [length]datatype{values} // An array
		myslice := myarray[start:end] // A slice made from the array
	*/

	// Samilar with array but no need to define size
	// like array but can change lenge; can append
	// cap() capacity of the slice (the number of elements the slice can grow or shrink to)
	friends := [] string {"Nutt", "Preaw"}

	// Append to Slice
	// slice = append(slice, element1, element2, ...)
	// slice3 = append(slice1, slice2...)
	// [https://www.w3schools.com/go/go_slices_modify.php]

	friends = append(friends, "Ice", "Ja")

	// copy slice
	// copy(dest, src) [https://www.w3schools.com/go/go_slices_modify.php]
	
	friendsNow := make([]string, len(friends), len(friends))
	copy(friendsNow, friends)
	fmt.Println(friendsNow)

	// Can access some element
	fmt.Println(friendsNow[1:3])

	// If-else
	if authorAge < 10 {
		fmt.Println("Young")
	} else if authorAge < 20 {
		fmt.Println("Teen")
	} else {
		fmt.Println("Adult")
	}

	// Switch 
	// No need break
	switch authorAge {
		case 20 : 
			fmt.Println("20")
		case 21 : 
			fmt.Println(21)
		default: 
			fmt.Println("> 21")
	}

	switch authorAge {
		case 19, 20 : 
			fmt.Println("19, 20")
		case 21, 22 : 
			fmt.Println(21, 22)
		default: 
			fmt.Println("> 21")
	}

	// Loop
	// continue : skip 1 time
	// break : stop loop
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	for index, value := range pets {
		fmt.Println(index, value)
	}

	fmt.Println("Check Animal:", checkAnimal(pets[:], amountPet, numPetCheck[:]))
	fmt.Println("            :",numPetCheck)
	fmt.Println("Sum of pets (real):", sumAnimal(numPetCheck[:]))
	isAll, notComplete := isEqualAnm(pets[:], amountPet[:], numPetCheck[:])
	fmt.Println("Complete:", isAll, "Missing:",notComplete)
}

// Function
func checkAnimal(pets [] string,  amountPet [] int, numPetCheck [] int) string {
	// Input
	// fmt.Scanf(stringFormat, address)
	for i := 0; i < len(pets); i++ {
		fmt.Printf("%v (All : %v) : ", pets[i], amountPet[i])
		fmt.Scanf("%d\n", &numPetCheck[i])
	}

	return "Finished";
}

func sumAnimal(numPet [] int) (result int) {
	for _, val := range numPet {
		result += val
	}
	return
}

func isEqualAnm(pets [] string,  amountPet [] int, numPetCheck [] int) (result bool, notComplete [] string){
	result = true
	for idx, val := range pets {
		if (numPetCheck[idx] != amountPet[idx]) {
			result = false
			notComplete = append(notComplete, val)
		}
	}
	return
}
