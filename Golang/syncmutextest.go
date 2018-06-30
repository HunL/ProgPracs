package main 

import ( 
    "fmt" 
    "sync" 
    "time" 
) 

var m *sync.Mutex 

func main() { 
    m = new(sync.Mutex) 
    go read(1) 
    go read(2) 
    time.Sleep(time.Second) 
} 

func read(i int) { 
    fmt.Println(i, "begin lock") 
    m.Lock() 
    fmt.Println(i, "in lock") 
    m.Unlock() 
    fmt.Println(i, "unlock") 
}

