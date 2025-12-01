// Package acdc provides a function to sum integers.
package acdc

// Sum takes a variable number of integer arguments and returns their sum.
func Sum(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}