package service

import (
	"sync"
	"time"

	"github.com/prime-factors/interfaces"
	"github.com/prime-factors/model"
)

type WithGoroutine struct {
	Prime interfaces.PrimeInterface
}

var wg sync.WaitGroup

// checking primes without using goroutine
func (w *WithGoroutine) PrimesCheck(numbers []int) model.PrimesChecking {
	t1 := time.Now()
	ch := make(chan model.Number, len(numbers))

	for _, number := range numbers {
		wg.Add(1)
		go w.primeCheck(ch, number)
	}
	wg.Wait()
	close(ch)

	t2 := time.Now()
	return model.PrimesChecking{
		Numbers:  numbers,
		Results:  toSlice(ch),
		Duration: w.Prime.DiffTime(t1, t2),
	}
}

func (w *WithGoroutine) primeCheck(ch chan model.Number, number int) {
	defer wg.Done()

	t1 := time.Now()
	isPrime := w.Prime.Check(number)
	factors := w.Prime.Factors(number)
	t2 := time.Now()

	ch <- model.Number{
		Digit:    number,
		Factors:  factors,
		IsPrime:  isPrime,
		Duration: w.Prime.DiffTime(t1, t2),
	}
}

func toSlice(c chan model.Number) []model.Number {
	s := make([]model.Number, 0)
	for i := range c {
		s = append(s, i)
	}
	return s
}
