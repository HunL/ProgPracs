package main

import (
    "fmt"
    "math/rand"
    "time"
)

func randomExchange(array []int64) {
    exchangeCounts := 4 * len(array)
    fmt.Println("counts: ", exchangeCounts)
    for i := 0; i < exchangeCounts; i++ {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        from := int64(r.Intn(len(array)))
        to := int64(r.Intn(len(array)))
        var tmp int64
        tmp = array[from]
        array[from] = array[to]
        array[to] = tmp
    }
}

func main() {
    array := []int64{1,2,3,4,5,6,7,8,9,10}
    fmt.Println(array)
    randomExchange(array)
    fmt.Println(array)
}
