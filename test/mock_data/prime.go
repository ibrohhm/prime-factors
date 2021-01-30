package mock_data

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type MockedPrime struct {
	mock.Mock
}

func (m *MockedPrime) Check(number int) bool {
	args := m.Called(number)
	return args.Bool(0)
}

func (m *MockedPrime) Factors(number int) []int {
	args := m.Called(number)
	return args.Get(0).([]int)
}

func (m *MockedPrime) DiffTime(t1 time.Time, t2 time.Time) time.Duration {
	args := m.Called(t1, t2)
	return args.Get(0).(time.Duration)
}
