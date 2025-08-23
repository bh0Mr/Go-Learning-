package main

import (
	"fmt"
	"strconv"
)

func main() {

	pracString()

}

func pracString(){


	num := 1223

	str1 := fmt.Sprint(num)

	fmt.Println(str1)

	fmt.Print(strconv.Itoa(num))

}