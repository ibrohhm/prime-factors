package interfaces

import "time"

type PrimeInterface interface {
	Check(number int) bool
	Factors(number int) []int
	DiffTime(t1 time.Time, t2 time.Time) time.Duration
}
