package service

import (
	"sync"
	"time"

	"github.com/prime-factors/model"
	"github.com/prime-factors/prime"
)

var wg sync.WaitGroup

// checking primes without using goroutine
func PrimesCheckWithGoroutine(numbers []int) model.PrimesChecking {
	t1 := time.Now()
	ch := make(chan model.Number, len(numbers))

	for _, number := range numbers {
		wg.Add(1)
		go numberCheckWithGoroutine(ch, number)
	}
	wg.Wait()
	close(ch)

	t2 := time.Now()
	return model.PrimesChecking{
		Numbers:  numbers,
		Results:  toSlice(ch),
		Duration: t2.Sub(t1),
	}
}

func numberCheckWithGoroutine(ch chan model.Number, number int) {
	defer wg.Done()

	t1 := time.Now()
	isPrime := prime.Check(number)
	factors := prime.Factors(number)
	t2 := time.Now()

	ch <- model.Number{
		Digit:    number,
		Factors:  factors,
		IsPrime:  isPrime,
		Duration: t2.Sub(t1),
	}
}

func toSlice(c chan model.Number) []model.Number {
	s := make([]model.Number, 0)
	for i := range c {
		s = append(s, i)
	}
	return s
}
