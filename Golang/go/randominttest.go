package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main(){
//    i := genRandNums()
//    fmt.Println("nums:",i,"len:",len(i))
    nums := genUnSame()
    fmt.Println("nums:",nums,"len:",len(nums))
    nth := nums[0:8]
    fmt.Println("nth:", nth, len(nth))

//    rand.Seed(time.Now().Unix())
    nums1 := genUnSame()
    fmt.Println("nums1:",nums1,"len1:",len(nums1))
    nth1 := nums[0:8]
    fmt.Println("nth1:", nth1, len(nth1))
}

func genRandNums() []int{
    rand.Seed(time.Now().Unix())

    randNums := make([]int, 0)
    fmt.Println("randNums:",randNums)

    for i := 0; i < 10; i++ {
        num := rand.Intn(10)
        randNums = append(randNums, num)
    }
    fmt.Println("nums:",randNums,"\n")

    return randNums
}


func genUnSame() []int {
    rand.New(rand.NewSource(time.Now().Unix()))

    randNums := make([]int, 0)
    fmt.Println("randNums:",randNums)

    unSame := rand.Perm(16)
    fmt.Println("unSame:",unSame)

    return unSame
}
