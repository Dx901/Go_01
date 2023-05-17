package main

import "fmt"

func main() {

	age :=35
	name := "Mkuru"

	//Print
	fmt.Print("Hello, ")
	fmt.Print("world \n")
	fmt.Print("new line \n")

	//println
	fmt.Println("Hello ninjas")
	fmt.Println("Adios pompeii")
	fmt.Println("My age is", age, "and my name is", name )

	//Formstted strings %_ = format specifier
	// /printf?
	fmt.Printf("may age is %v and my name is %v \n", age, name)
	fmt.Printf("may age is %q and my name is %q \n", age, name)
	fmt.Printf("age is f type %T", age)
	fmt.Printf("Yu x=score %0.2f points \n", 222.6678778)

	// Sprint f(Saces the string in a variable that can be used 
	// somewhere else) (save formatted string)
		var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
		fmt.Println("The saved string is: ", str)


}