package main

import (
	"fmt"
)

func main() {
	fmt.Println(subString("สกลนคร"))
}

func mainOfMap() {
	var m map[int]int
	m = make(map[int]int)
	m[0] = 1
	m[1] = 2
	m[2] = 3

	for k, v := range m {
		fmt.Println(k, v)
	}

	s := []int{1, 2, 3}
	for i, v := range s {
		fmt.Println(i, v)
	}

}

func count(i int) int {
	n := 0
	for i := 0; i < i; i++ {
		n += i
	}
	return n
}

type pFunc func() int

func one() int   { return 1 }
func Peter() int { return one() + one() }

func factory(a, b pFunc) pFunc {
	return func() int {
		return a() + b()
	}
}

func addp(p *int, n int) {
	*p += n
}

func subString(s string) []string {
	s += "*"
	couple := []string{}
	for len(s) > 1 {
		couple = append(couple, s[:2])
		s = s[2:]
	}
	return couple
}

func reverse(four [4]int) [4]int {
	const lenght = 4
	var revered [4]int
	for i, reversedIndex := 0, lenght-1; i < lenght; i, reversedIndex = i+1, reversedIndex-1 {
		revered[reversedIndex] = four[i]

	}
	return revered
}

func isEven(n int) bool {
	return n%2 == 0
}

func printEven(from, upperBond int) {
	for i := from; i < upperBond; i++ {
		if isEven(i) {
			fmt.Print(i, " ")
		}
	}
}

func printPrime(n int) {
	max := n + 1
	for i := 2; i <= max; i++ {
		count := 0
		for j := 2; j <= max; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 1 {
			fmt.Print(i, " ")
		}
	}
}

func monthlyIntallment(financeAmount float64, rate float64, year float64) float64 {
	const totalMonthInYear = 12
	annualInterestYearRate := rate / 100
	installment := totalMonthInYear * year
	return ((financeAmount * annualInterestYearRate * year) + financeAmount) / installment
}

func add[T int | float64](a, b T) T {
	return a + b
}
