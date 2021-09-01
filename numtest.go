package main

import (
    "fmt"
    "time"
)

func main() {
    var myint int
    var limit int
    var threemultiples int

    fmt.Println("Welcome to the Go number test")
    start := time.Now().Unix()
    fmt.Println("Enter a number to count to")
    fmt.Scanln(&limit)

    for i := 0; i < limit; i++ {
        fmt.Printf("Current Number: %d\n", myint)
        if myint % 3 == 0{
            threemultiples+=1
        }
        if myint % 2 == 0{
            fmt.Println("Number is even\n")
        }else {
            fmt.Println("Number is odd\n")
        }
        myint += 1
    }

    end := time.Now().Unix()
    fmt.Println(myint) // prints out limit
    fmt.Printf("Execution time: %d seconds\n", end - start)
    fmt.Printf("Total 3 multiples: %d\n", threemultiples)
}
