package main

import "fmt"

func main() {
	var blah string
	blah = "Hello"

	fmt.Println(blah)

	const c string = "immutable"
	fmt.Println(c)

	// c = "mutated"

	// Syntax sugar
	i := 0

	var j int = 2

	fmt.Printf("j is: %d\n", i)
	fmt.Printf("j is: %d\n", j)

	// Arrays

	var arr []string

	arr = append(arr, "vic")
	arr = append(arr, "patrick")
	arr = append(arr, "stephanie")

	fmt.Println(arr)

	for index, item := range arr {
		fmt.Printf("Index: %d Item: %s\n", index, item)
	}

	a, b := tuple()

	fmt.Println(a)
	fmt.Println(b)

	sum := 0
	for x := 0; x < 10; x++ {
		sum += x
	}

	fmt.Println(sum)

	array := [...]float64{7.0, 8.5, 9.1}
	z := Sum(&array) // Note the explicit address-of operator

	fmt.Println(z)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

func tuple() (string, string) {
	return "hello", "world"
}
