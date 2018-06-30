package main

import (
    . "fmt"
    "sort"
)

type User struct {
    OpenID string
    Name string
    Grade int
}

type Users []User

func (u Users) Len() int {
    return len(u)
}

func (u Users) Less(i, j int) bool {
    return u[i].Grade > u[j].Grade
}

func (u Users) Swap(i, j int) {
    u[i], u[j] = u[j], u[i]
}

func main() {
    users := Users{
        {
            OpenID: "test1",
            Name: "test1",
            Grade: 90,
        },
        {
            OpenID: "test2",
            Name: "test2",
            Grade: 80,
        },
        {
            OpenID: "test3",
            Name: "test3",
            Grade: 70,
        },
        {
            OpenID: "test4",
            Name: "test4",
            Grade: 90,
        },
        {
            OpenID: "test5",
            Name: "test5",
            Grade: 100,
        },
    }

    Println("before sort")
    for _, user := range users {
        Println(user.Name, ":", user.Grade)
    }

    sort.Sort(users)

    Println("\nafter sort")
    for _, user := range users {
        Println(user.Name, ":", user.Grade)
    }

    Println("\nlen: ", users.Len())
}
