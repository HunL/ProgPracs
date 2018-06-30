package main

import . "fmt"
import "reflect"

func main() {
    var circle float64 = 6.28
    var icir interface{}

    icir = circle
    Println("Reflect : circle.Value = ", reflect.ValueOf(icir))
    Println("Reflect : circle.Type = ", reflect.TypeOf(icir))

    valueref := reflect.ValueOf(icir)
    Println(valueref)
    Println(valueref.Interface())

    y := valueref.Interface().(float64)
    Println(y)

}
