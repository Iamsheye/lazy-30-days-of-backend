package main

import "fmt"

func main() {
	fmt.Println(LeapYear(2195, 3))
	fmt.Println(LeapYear(2799, 10))
}

func LeapYear(year int, num int) []int {
	leapYears := []int{}

	for i := year + 1; len(leapYears) < num; i++ {
		if i%4 == 0 && i%100 != 0 || i%400 == 0 {
			leapYears = append(leapYears, i)
		}
	}

	return leapYears
}
