package fib

import (
    "testing"
    . "fmt"
    "runtime"
)

func TestFib(t *testing.T) {
    var testdatas = []struct {
        n    int
        want int64
    }{
        {0, 0},
        {1, 1},
        {2, 1},
        {3, 2},
        {4, 3},
        {16, 987},
        {32, 2178309},
        {45, 1134903170},
    }

    for _, test := range testdatas {
        n := test.n
        want := test.want
        got := fib1(n)
        // got := fib2(n)
        // got := fib3(n)
        // got := fib4(n)

        if got != want {
            t.Errorf("fib(%d)=%d, want %d\n", n, got, want)
        }
    }
}

// recursive version
func fib1(n int) int64 {
    // Println("recursive11111--------------->>>> version, n: ", n)
    if n == 0 || n == 1 {
        return int64(n)
    }
    return fib1(n-1) + fib1(n-2)
}

// concurrent version
func fib2(n int) int64 {
    Println("concurrent version")
    var got int64
    var channel = make(chan int64, 2)

    if n == 0 || n == 1 {
        return int64(n)
    }

    runtime.GOMAXPROCS(2)

    go func() { channel <- fib1(n - 2) }()
    go func() { channel <- fib1(n - 1) }()

    got = <-channel
    Println("got1: ", got)
    got += <-channel
    Println("got2: ", got)
    return got
}

// iteration version
func fib3(n int) int64 {
    Println("iteration version")
    var a, b int64
    a, b = 0, 1

    for i := 0; i < n; i++ {
        a, b = b, a+b
    }
    return a
}

// wrapper version
func fibWrapper4() func() int64 {
    var a, b int64
    a, b = 0, 1

    return func() int64 {
        a, b = b, a+b
        return a
    }
}

func fib4(n int) int64 {
    Println("wrapper version")
    var got int64
    got = 0
    f := fibWrapper4()
    for i := 0; i < n; i++ {
        got = f()
    }
    return got
}
