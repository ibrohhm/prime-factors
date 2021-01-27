package request

import (
	"errors"
)

type Form struct {
	ExampleType string `json:"example_type"`
	Numbers     []int  `json:"numbers"`
	IsGoroutine bool   `json:"is_goroutine"`
}

func (self *Form) Validate() error {
	if self.ExampleType == "small" || self.ExampleType == "medium" || self.ExampleType == "big" || self.ExampleType == "custom" {
		return nil
	}

	return errors.New("example_type must be small, medium, big or custom")
}
