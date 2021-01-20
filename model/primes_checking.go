package model

import "time"

type PrimesChecking struct {
	Numbers  []int
	Results  []Number
	Duration time.Duration
}
