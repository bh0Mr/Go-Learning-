package main

import "fmt"




func main() {

	Boxes()
   
}

func Boxes(){



var boxes [3]int

    boxes[0] = 10
    boxes[1] = 20
    boxes[2] = 30

    fmt.Println("Box 1 has:", boxes[0])
    fmt.Println("Box 2 has:", boxes[1])
    fmt.Println("Box 3 has:", boxes[2])

    fmt.Println("All boxes:")
    for i := 0; i < len(boxes); i++ {
        fmt.Printf("Box %d has: %d\n", i+1, boxes[i])
    }

    boxes[1] = 99
    fmt.Println("After modifying Box 2:", boxes[1])

}
