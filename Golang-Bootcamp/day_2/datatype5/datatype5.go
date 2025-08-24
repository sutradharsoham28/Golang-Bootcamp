// Go program to illustrate
// the use of booleans
package main

import "fmt"

func main() {

	str1 := "Soham"
	str2 := "Sutradhar"
	str3 := "Soham"
	var str4 string = "Golang"

	result1 := str1 == str2
	result2 := str1 == str3

	fmt.Println(result1)
	fmt.Println(result2)

	fmt.Printf("The type of result1 is %T and the type of result2 is %T", result1, result2)
	fmt.Printf("\nLength of the string %s is:%d", str1, len(str1))
	fmt.Println("\nNew string : ", str1+str2) //String concatenation
	fmt.Print(str4)

}
