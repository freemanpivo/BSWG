package main

import "fmt"

type DayOfTheWeek uint8

const (
	Monday DayOfTheWeek = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func ops(a int, b int) (int, int) {
	return a + b, a - b
}

// funcao variadica
func sum_variadic(nums ...int) int {
	total := 0
	for _, a := range nums {
		total = total + a
	}

	return total
}

func main() {
	fmt.Printf("Monday is %d\n", Monday)
	fmt.Printf("Wednesday is %d\n", Wednesday)
	fmt.Printf("Friday is %d\n", Friday)

	sum, sub := ops(2, 2)
	fmt.Println("2+2 = ", sum, "2-2 = ", sub)

	sum2, _ := ops(10, 2)
	fmt.Println("10+2 = ", sum2)

	total := sum_variadic(1, 2, 3, 4, 5)
	fmt.Println("Soma de 1 a 5 = ", total)
}
