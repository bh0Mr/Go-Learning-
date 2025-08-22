package main

import "fmt"

func main() {


    str := "hello, world"

    fmt.Println("Length of string:", len(str))
    fmt.Println("First byte:", str[0])        

    fmt.Println("Substring:", str[0:5])
    fmt.Println("Substring:", str[7:])

    str += ", Golang!"
    fmt.Println("After concatenation:", str)

    //imutable
<<<<<<< HEAD
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	
=======

>>>>>>> Stashed changes
=======

>>>>>>> Stashed changes
=======

>>>>>>> ak/feature
    imu := str
    imu += " Great!"
    fmt.Println("Original string:", str)
    fmt.Println("New string:", imu)

    raw := `This is just a Practice for go language 
multile line string \n`
    fmt.Println("Raw string:", raw)

<<<<<<< HEAD
<<<<<<< Updated upstream
<<<<<<< Updated upstream
    escape := "Hello, \"world\"!"
    fmt.Println("String with escape sequences:", escape)
=======
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
=======
>>>>>>> ak/feature
}
