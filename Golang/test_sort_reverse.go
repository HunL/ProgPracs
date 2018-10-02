package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type PerColl []Person

var persons PerColl

func (PerColl) Len() int {
	return len(persons)
}

func (PerColl) Less(i, j int) bool {
	// return persons[i].Name < persons[j].Name
	return persons[i].Age < persons[j].Age
}

func (PerColl) Swap(i, j int) {
	persons[i], persons[j] = persons[j], persons[i]
}

func main() {

	persons = []Person{
		Person{"Cynhard", 5},
		Person{"Gopher", 16},
		Person{"Monkey", 12},
	}

	sort.Sort(persons)
	fmt.Println(persons)                // [{Cynhard 5} {Gopher 11} {Monkey 12}]
	fmt.Println(sort.IsSorted(persons)) // true
	sort.Sort(sort.Reverse(persons))
	fmt.Println(persons) // [{Monkey 12} {Gopher 11} {Cynhard 5}]
}
