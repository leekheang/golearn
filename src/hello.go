package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hello World")

	var a = 5
	fmt.Println(a)

	var b [5]int
	b[2] = 7
	fmt.Println(b)

	c := []int{5, 4, 3, 2, 1}
	c = append(c, 13)
	fmt.Println(c)

	for d := 0; d < 5; d++ {
		fmt.Println(d)
	}

	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value", value)
	}

	m := make(map[string]string)
	m["x"] = "lee"
	m["y"] = "kheang"

	for key, value := range m {
		fmt.Println("Key:", key, "value:", value)
	}

	result := sum(2, 3)
	fmt.Println(result)

	resultsqrt, err := sqrt(64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultsqrt)
	}

	personone := person{name: "Lee", age: 23}
	fmt.Println("My name is :", personone.name, "I'm ", personone.age)

	r := 7
	fmt.Println("add & make it to pointer:", &r)
}

func sum(n int, o int) int {
	return n + n
}

func sqrt(p float64) (float64, error) {
	if p < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(p), nil
}
