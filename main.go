package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("############################ Declaring Variables & Place Holders in Go ###############################")
	var greet = "Hello, World!"
	greet2 := "Hello, Go!"
	myinteger := 42
	myfloat64 := 3.14
	myboolean := true
	myarray := [...]int{1, 2, 3, 4, 5}
	myslice := []string{"Go", "is", "fun"}
	mapvar := map[string]int{"go": 1, "python": 2, "java": 3}
	myconst := "This is a constant"

	fmt.Printf("String: %s\nQuoted String: %q\nInteger: %d\nFloat: %.2f\nBoolean: %t\nArray: %v\nSlice: %v\nMap: %v\nConstant: %v\n",
		greet, greet2, myinteger, myfloat64, myboolean, myarray, myslice, mapvar, myconst)

	fmt.Println("############################ Inner and Outer blocks in Go ###############################")
	fmt.Println("")
	{
		InnerBlock := "This is the inner block"
		fmt.Printf("greet: %s \nInnerBlock: %s", greet, InnerBlock)
	}
	fmt.Println("\nThis is the outer block")
	// fmt.Printf("Error: InnerBlock is not accessible here %s\n",InnerBlock) // This will cause an error because InnerBlock is not defined in this scope

	fmt.Println("############################ Input From User in Go ###############################")
	var user_greet string //we must declare them first before using them
	var user_boolean bool
	var user_integer int
	// fmt.Print("Enter greeting: ")
	// fmt.Scan(&user_greet) // Read string

	// fmt.Print("Enter boolean (true/false): ")
	// fmt.Scan(&user_boolean) // Read boolean

	// fmt.Print("Enter integer: ")
	// fmt.Scan(&user_integer)

	// fmt.Println("Error and count at scanf")
	// count, err := fmt.Scanf("%s %t %d", &user_greet, &user_boolean, &user_integer)
	// fmt.Println("Count: ", count)
	// fmt.Println("Error: ", err)

	fmt.Println("############################ Data Conversion in Go ###############################")
	user_greet = "Hello"
	user_boolean = true
	user_integer = 100
	fmt.Printf("Integer to Float64: %.2f\n", float64(user_integer))
	fmt.Printf("Integer to string: %s\n", strconv.Itoa(user_integer))
	i, err := strconv.Atoi(user_greet)
	fmt.Printf("String to Integer: %d\n Error : %v", i, err)
	fmt.Printf("Boolean to string: %s\n", strconv.FormatBool(user_boolean))
	fmt.Println("############################ If and Switch Statement in Go ###############################")
	if user_integer < 50 {
		fmt.Println("The integer is Less than 50")
	} else if user_integer == 50 {
		fmt.Println("The integer is equal to 50")
	} else {
		fmt.Println("The integer is Greater than 50")
	}

	switch {
	case user_integer > 50:
		fmt.Println("Switch: The integer is Greated than 50")
		fallthrough
	case user_integer == 50:
		fmt.Println("Switch: The integer is equal to 50")
		fallthrough
	default:
		fmt.Println("Switch: The integer is less than 50")
	}
	fmt.Println("############################ For loop Statement in Go ###############################")
	for i := 0; i < 5; i++ {
		fmt.Printf("For Loop Iteration: %d\n", i)
		if i == 2 {
			break
		}
	}
	fmt.Println("############################ Array and Slices in Go ###############################")
	var myarray2 = [2]int{10, 20}
	myarray3 := [...]int{30, 40, 50}
	for i := 0; i < len(myarray2); i++ {
		fmt.Printf("Array Element at index %d: %d\n", i, myarray2[i])
	}
	for index, value := range myarray3 {
		fmt.Printf("Array Element at index %d: %d\n", index, value)
	}
	multi_dimensionalArray := [2][3]int{{1, 2, 5}, {3, 4, 5}}
	fmt.Printf("%v\n", multi_dimensionalArray)
	multi_dimensionalArray[0][1] = 10
	fmt.Printf("After modification: %v\n", multi_dimensionalArray)

	var myslice2 = []int{1, 2, 3, 4, 5}
	myslice3 := []int{6, 7, 8, 9, 10}
	fmt.Printf("Full array :%v\n", myarray2)
	slice_of_array := myarray2[0:2]
	slice_of_array[0] = 100
	fmt.Printf("Full array after modification :%v\n", myarray2)
	myslice4 := make([]int, 0, 10) // Length 0, capacity 10 to avoid reallocation
	myslice4 = append(myslice4, 6)
	myslice4 = append(myslice4, myslice2...)
	myslice5 := make([]int, 0, 10)
	number_of_elements := copy(myslice5, myslice3)
	fmt.Printf("Slice of array: %v\n", number_of_elements)
	fmt.Println("############################ Maps in Go ###############################")
	var map_name map[string]int = map[string]int{"apple": 1, "banana": 2, "cherry": 3}
	code := map[string]int{"go": 1, "python": 2, "java": 3}
	fmt.Println(map_name["apple"])
	delete(code, "banana")
	fmt.Println("Map after deletion:", code)
	fmt.Println("############################ Functions in Go ###############################")
	add_result,add_result_int:= add (5, 10)
	fmt.Println(add_result)
	fmt.Println(add_result_int)
	// defer Print(1, "Hello", "World", "Go is great!") // defer in functions
	fmt.Println(factorial(5)) // Output: 120
	x:=func (a int) int {
		return a * a
	}
	x(5) // This is an anonymous function that squares the input
	fmt.Println("Square of 5 is:", x(5))
	fmt.Println(applyoperation(3, 5, add2))      // Output: 8
	fmt.Println(applyoperation(3, 5, multiply2)) // Output: 15

	fmt.Println("############################ Pointers in Go ###############################")
	Pointer_var:= 77
	// ptr:=*x
	fmt.Println(Pointer_var,&Pointer_var,*(&Pointer_var))
	ptr:=&Pointer_var
	*ptr=88
	fmt.Println(Pointer_var,&Pointer_var,*(&Pointer_var))
	myinteger=10
	modify(myinteger)
	fmt.Println("pass my value Pointer:", myinteger) 
	modify2(&myinteger)
	fmt.Println("pass my reference Pointer:", myinteger) 
	
	fmt.Println("pass my reference Pointer at slice:", add_to_slice(myslice)) 
	fmt.Println("############################ struct in Go ###############################")
	mystruct1:=mystruct{name: "John", age: 30}
	mystruct2 := mystruct{"Doe", 25}
	fmt.Printf("Struct Name: %+v\nStruct Age: %+v\n", mystruct1.name, mystruct1.age)
	if mystruct1.age > mystruct2.age {
		fmt.Println("To Compare we must se the sa")
	}
}

func add(a int,b int) (string,int) {
		sum :=a+b
		return "Sum is: " + strconv.Itoa(sum), sum
	}
// Function with variadic parameters
func Print(a int, s ...string) (int, int) {
	for _, str := range s {
		fmt.Println(str)
	}
	return 0, 0
}

// Recursive function to calculate factorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

}
// high order function
func applyoperation(a,b int ,operation func(int, int) int) int {
	return operation(a,b)
}

func add2(x, y int) int {
 return x + y
}

func multiply2(x, y int) int {
 return x * y
}

func modify(a int){
	a+= 10
}

func modify2(a *int){
	*a+= 10
}

func add_to_slice(s []string)[]string{
	s=append(s,"and","easy")
	return s
}

type mystruct struct {
	name string
	age  int
}