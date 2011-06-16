package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(CalculatePi(300000))
}


func CalculatePi(term int) float64 {
    
    ch := make(chan float64)

    //map
    for i := 0; i <= term; i++ {
        go formula(ch, float64(i))
    }
    result := 0.0 
    //reduce
    for i := 0; i <= term; i++ {
        result += <- ch
    }
    return result
}


func formula(ch chan float64, term float64) {
    result := 4 * (math.Pow(-1, term)/(2.0 * term + 1.0))
    ch <- result
}


