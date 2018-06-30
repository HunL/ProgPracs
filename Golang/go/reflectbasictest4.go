package main

import . "fmt"
import "reflect"

func main() {
    var circle float64 = 6.28

    value := reflect.ValueOf(circle)
    Println("Reflect : value = ", value)
    value.SetFloat(3.14)
}
