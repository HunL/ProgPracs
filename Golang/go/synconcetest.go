package main 

import ( 
    "fmt" 
    "sync" 
    "time" 
) 

func main() { 
    var once sync.Once 
    for i := 0; i < 20; i++ { 
        go func() { 
            once.Do(read) 
        }() 
    } 
    time.Sleep(time.Second) 
} 

func read() { 
    fmt.Println(1) 
}
