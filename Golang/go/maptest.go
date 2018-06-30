package main

import "fmt"

func main(){
    monthdays := map[string]int{
        "Jan":31, "Feb":28, "Mar":31,
        "Apr":30, "May":31, "Jun":30,
        "Jul":31, "Aug":31, "Sep":30,
        "Oct":31, "Nov":30, "Dec":31,
    }

    fmt.Printf("%d\n", monthdays["Jan"])


    year := 0
    for _, days := range monthdays {
        year += days
    }
    fmt.Printf("Numbers of days in a year: %d\n", year)


    v,ok := monthdays["Mar"]
    fmt.Println(v, ok)

    v1,ok1 := monthdays["undefined"]
    fmt.Println(v1, ok1)

    delete(monthdays, "Mar")
    v2,ok2 := monthdays["Mar"]
    fmt.Println(v2, ok2)

    monthdays["undefined"] = 40
    v3,ok3 := monthdays["undefined"]
    fmt.Println(v3, ok3)
}
