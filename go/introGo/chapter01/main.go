/*Create a new executable program that references the fmt library and contains one function called main. That function takes no arguments and doesn’t return anything. It accesses the Println function contained inside of the fmt package and invokes it using one argument—the string Hello, World.*/

// package declaration, every go program must start with it
package main

// fmt is the format package for io
import "fmt"

// this is a comment

/* and so is this
line */

func main() {
	fmt.Println("Hello, world")
}
