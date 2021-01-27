package model

import "time"

type Number struct {
	Digit    int           `json:"digit"`
	Factors  []int         `json:"factors"`
	IsPrime  bool          `json:"is_prime"`
	Duration time.Duration `json:"duration"`
}
