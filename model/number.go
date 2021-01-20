package model

import "time"

type Number struct {
	Digit    int
	Factors  []int
	IsPrime  bool
	Duration time.Duration
}
