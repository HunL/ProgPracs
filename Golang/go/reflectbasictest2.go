package main

import . "fmt"
import "reflect"

func main() {
    var circle float64 = 6.28

    Println("Reflect : circle.Value = ", reflect.ValueOf(circle))
    Println("Reflect : circle.Type = ", reflect.TypeOf(circle))
}
