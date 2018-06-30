package main

import (
    "fmt"
    "sync"
)

var l sync.Mutex

func main(){
  a := 1;
  p_a := &a;
  var wg sync.WaitGroup
  wg.Add(2)
  go func(p_a *int){
      l.Lock()
    for i := 1; i < 100000; i++ {
      *p_a += 1
    }
      l.Unlock()
    wg.Done()
  }(p_a)
  go func(p_a *int){
      l.Lock()
    for i := 1; i < 100000; i++ {
      *p_a -= 1
    }
      l.Unlock()
    wg.Done()
  }(p_a)
  wg.Wait()
  fmt.Println(*p_a)
}
