package main

import "fmt"

type CardInfo struct{
    X uint8
    Y uint8
}

type User struct{
    id uint8
    name string
    card []*CardInfo
}

func main() {
    card := []*CardInfo{&CardInfo{X:1, Y:1}, &CardInfo{X:2, Y:2}}
    user := &User{id:1, name:"aaa", card:card}
    fmt.Printf("user:%s\n", user)


    card1 := &CardInfo{X:1, Y:2}
    fmt.Printf("\ncard: %s\n", card)
    card = append(card, card1)
    fmt.Printf("new card: %s\n", card)


    samecard := &CardInfo{X:1, Y:1}
    card = delete(card, samecard)
    fmt.Printf("delete card: %s\n", card)

    user = &User{id:1, name:"bbb"}
    fmt.Println("user:", user)
}
