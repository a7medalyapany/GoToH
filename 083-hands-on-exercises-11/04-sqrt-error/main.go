package main

import (
	"fmt"
)

type sqrtError struct {
	lat  string
	long  string
	err  error
}

func (e *sqrtError) Error() string {
	return fmt.Sprintf("Math error at lat: %s long: %s: %v", e.lat, e.long, e.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		fmt.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, &sqrtError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  fmt.Errorf("cannot Sqrt negative number: %v", f),
		}
	}

	return 42, nil
}