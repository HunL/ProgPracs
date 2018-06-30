package main

import . "fmt"
import "reflect"

func main() {
    var circle float64 = 6.28

    value := reflect.ValueOf(circle)
    Println("Reflect : value = ", value)
    //value.SetFloat(3.14)
    Println("Settability of value: ", value.CanSet())

    value2 := reflect.ValueOf(&circle)
    Println("Settability of value2: ", value2.CanSet())

    value3 := value2.Elem()
    Println("Settability of value3: ", value2.CanSet())

    value3.SetFloat(3.14)
    Println("Value of value3: ", value3)
    Println("Value of circle: ", circle)
}
