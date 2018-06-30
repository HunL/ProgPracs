package main

import (
    . "fmt"
    "time"
)

func main() {
    t1 := time.NewTicker(1 * time.Second)
    num := 0
    for {
        if num > 5 {
            //t1.Stop();
            break;
        }

        select {
            case <-t1.C:
                num++;
                Println("1 second...");
        }
    }


    Println("still print")
    Println("t1.C: ", t1.C)

    for k := range t1.C {
        Println("Tick at: ", k)
    }
}
