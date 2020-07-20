// Package leap do the calculations to determine if a given year
// year is leap or not.
package leap

import "fmt"

// Receive a year as an int and return a boolean if is leap (true) or not (false)
func IsLeapYear(year int) bool {
	if (year%4 == 0) && (year%100 != 0) {
		fmt.Printf("The year %v is leap year\n", year)
		return true
	} else if (year%100 == 0) && (year%400 == 0) {
		fmt.Printf("The year %v is leap year\n", year)
		return true
	}
	fmt.Printf("The year %v is not a leap year\n", year)
	return false
}
