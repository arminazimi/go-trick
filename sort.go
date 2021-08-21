package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	s := []int{6, 3, 8, 6, 9, 4, 7, 2}
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)

	a := []string{"w", "c", "b", "d", "z", "f", "m", "a"}
	fmt.Println(a)
	sort.Strings(a)
	fmt.Println(a)

	peopleA := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(peopleA)
	sort.Sort(ByAge(peopleA))
	fmt.Println(peopleA)

	peopleN := []Person{
		{"z", 31},
		{"s", 42},
		{"l", 17},
		{"a", 26},
	}
	fmt.Println(peopleN)
	sort.Sort(ByName(peopleN))
	fmt.Println(peopleN)

}
