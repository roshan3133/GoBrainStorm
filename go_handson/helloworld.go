package main

import "fmt"

    //1. Define variables outside of functions
    var a string = "outside function variable"
    b := "outside function variable"

func main() {
    //2. Define type of variable like 
    var x string = "Hello World"
    var y int = 10
    fmt.Println("x", x, "y", y)

    //3. Without Defining type of variable
    x := "Hello World"
    y := 10
    fmt.Println("x", x, "y", y)
    fmt.Println("a", a, "b", b)
    }
