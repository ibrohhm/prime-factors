package service

import (
	"time"

	"github.com/prime-factors/model"
	"github.com/prime-factors/prime"
)

// checking primes without using goroutine
func PrimesCheckWithoutGoroutine(numbers []int) model.PrimesChecking {
	var results []model.Number

	t1 := time.Now()
	for _, number := range numbers {
		results = append(results, numberCheckWithoutGoroutine(number))
	}
	t2 := time.Now()

	return model.PrimesChecking{
		Numbers:  numbers,
		Results:  results,
		Duration: t2.Sub(t1),
	}
}

func numberCheckWithoutGoroutine(number int) model.Number {
	t1 := time.Now()
	isPrime := prime.Check(number)
	factors := prime.Factors(number)
	t2 := time.Now()

	return model.Number{
		Digit:    number,
		Factors:  factors,
		IsPrime:  isPrime,
		Duration: t2.Sub(t1),
	}
}
