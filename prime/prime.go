package prime

import (
	"math"
	"time"
)

type Prime struct {
}

// checking prime number using sieve method
func (p *Prime) Check(number int) bool {
	if number == 2 {
		return true
	} else if number%2 == 0 || number == 1 || number <= 0 {
		return false
	} else {
		sieve_max := sqrt(number)
		i := 3

		for i <= sieve_max {
			if number%i == 0 {
				return false
			}
			i += 2
		}
	}

	return true
}

// finding all prime factors
func (p *Prime) Factors(number int) []int {
	var factors []int
	if number == 1 {
		return []int{1}
	} else {
		n := number
		for i := 2; i <= number; i++ {
			if n%i == 0 {
				factors = append(factors, i)
			}
			// eliminate factor i on n, ie: n=36, i=2 -> return 9
			for n%i == 0 {
				n /= i
			}
		}
	}

	return factors
}

// duration from t1 to t2 (unit test purpose)
func (p *Prime) DiffTime(t1 time.Time, t2 time.Time) time.Duration {
	return t2.Sub(t1)
}

func sqrt(number int) int {
	f := float64(number)
	return int(math.Sqrt(f))
}
