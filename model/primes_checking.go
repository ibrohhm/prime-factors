package model

import "time"

type PrimesChecking struct {
	Numbers  []int         `json:"numbers"`
	Results  []Number      `json:"results"`
	Duration time.Duration `json:"duration"`
}
