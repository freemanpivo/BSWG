package main

import "fmt"

func main() {
	theSum := doit(sum, 2, 3)
	theMultiply := doit(multiply, 2, 3)

	fmt.Println("2+3 = ", theSum)
	fmt.Println("2*3 = ", theMultiply)

	ones := accumulator(1)
	twos := accumulator(2)

	fmt.Println("a", "b")
	for range 5 {
		fmt.Println(ones(), twos())
	}
}

func doit(operator func(int, int) int, a int, b int) int {
	return operator(a, b)
}

func sum(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func accumulator(increment int) func() int {
	i := 0
	return func() int {
		i = i + increment
		return i
	}
}
