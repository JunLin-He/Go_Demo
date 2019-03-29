package main

import (
    "fmt"
    "math"
)

func ConvertIntToInt16(x int) int16 {
    if math.MinInt16 <= x && x <= math.MaxInt16 {
        return int16(x)
    }

    panic(fmt.Sprintf("%d is out of int16 range", x)) 
}

func Int16FromInt(x int) (i int16, err error) {
    defer func() { // Delay the running of anonymous function, and pick up the panic by recover, and transform panic to error
        if e := recover(); e != nil {
            err = fmt.Errorf("%v", e)
        }
    }()
    i = ConvertIntToInt16(x)
    return i, nil
}

func main() {
    if _, e := Int16FromInt(65567); e != nil {
        fmt.Printf("%v\n", e)
    } else {
        fmt.Printf("no errors\n")
    }
}


