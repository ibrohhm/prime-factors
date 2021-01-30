package service

import (
	"time"

	"github.com/prime-factors/interfaces"
	"github.com/prime-factors/model"
)

type WithoutGoroutine struct {
	Prime interfaces.PrimeInterface
}

// checking primes without using goroutine
func (w *WithoutGoroutine) PrimesCheck(numbers []int) model.PrimesChecking {
	var results []model.Number

	t1 := time.Now()
	for _, number := range numbers {
		results = append(results, w.primeCheck(number))
	}
	t2 := time.Now()

	return model.PrimesChecking{
		Numbers:  numbers,
		Results:  results,
		Duration: w.Prime.DiffTime(t1, t2),
	}
}

func (w *WithoutGoroutine) primeCheck(number int) model.Number {
	t1 := time.Now()
	isPrime := w.Prime.Check(number)
	factors := w.Prime.Factors(number)
	t2 := time.Now()

	return model.Number{
		Digit:    number,
		Factors:  factors,
		IsPrime:  isPrime,
		Duration: w.Prime.DiffTime(t1, t2),
	}
}
