package main

import (
	"fmt"

	"github.com/repoofideas/gorganon/person"
)

func plus(a, b int, name string) (int, string) {
	return a + b, name
}

func sumAll(a ...int) int {
	var total int
	for _, item := range a {
		total += item
	}
	return total
}

// type person struct {
// 	name string
// 	age  int
// }

// func (p person) sayHello() {
// 	fmt.Printf("Hello! My name is %s and I'm %d", p.name, p.getKoreanAge())
// }

// func (p person) getKoreanAge() int {
// 	year := 1992
// 	return time.Now().Year() - year + 1
// }

func main() {
	// result, name := plus(2, 2, "kris")
	// fmt.Println(result, name)
	// fmt.Println(sumAll(1, 2, 3, 4, 5))
	// fmt.Println(result, fmt.Sprintf("%b\n", result))
	// // [int] for int size array
	// // [] for slice (equivalent to dynamic array)
	// foods := []string{"kimchi", "yogurt", "chicken"}
	// for _, dish := range foods {
	// 	fmt.Println(dish)
	// }
	// // appending in go
	// foods = append(foods, "curry")
	// for i := 0; i < len(foods); i++ {
	// 	fmt.Println(foods[i])
	// }

	// a := 2
	// b := &a
	// a = 50
	// fmt.Println(&a, b, *b)

	// kris := person{"kris", 29}
	// kris.sayHello()

	kris := person.Person{}
	kris.SetDetails("kris", 29)
	fmt.Println(kris)
}
