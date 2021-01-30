package service_test

import (
	"testing"
	"time"

	"github.com/prime-factors/model"
	"github.com/prime-factors/service"
	"github.com/prime-factors/test/mock_data"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	WithGoroutine    service.WithGoroutine
	WithoutGoroutine service.WithoutGoroutine
	Numbers          []int
	Response         model.PrimesChecking
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.Numbers = []int{3, 12}
	suite.Response = model.PrimesChecking{
		Numbers: suite.Numbers,
		Results: []model.Number{
			model.Number{
				Digit:    3,
				Factors:  []int{3},
				IsPrime:  true,
				Duration: time.Duration(1),
			},
			model.Number{
				Digit:    12,
				Factors:  []int{2, 3},
				IsPrime:  false,
				Duration: time.Duration(1),
			},
		},
		Duration: time.Duration(1),
	}

	mockedPrime := new(mock_data.MockedPrime)
	mockedPrime.On("Check", 12).Return(false)
	mockedPrime.On("Check", 3).Return(true)
	mockedPrime.On("Factors", 12).Return([]int{2, 3})
	mockedPrime.On("Factors", 3).Return([]int{3})
	mockedPrime.On("DiffTime", mock.Anything, mock.Anything).Return(time.Duration(1))
	suite.WithGoroutine = service.WithGoroutine{
		Prime: mockedPrime,
	}
	suite.WithoutGoroutine = service.WithoutGoroutine{
		Prime: mockedPrime,
	}
}
func (suite *ServiceTestSuite) TestPrimesWithoutGoroutineCheck() {
	suite.Equal(suite.WithoutGoroutine.PrimesCheck(suite.Numbers), suite.Response)
}

func (suite *ServiceTestSuite) TestPrimesWithGoroutineCheck() {
	resp := suite.WithGoroutine.PrimesCheck(suite.Numbers)
	suite.Equal(resp.Numbers, suite.Numbers)
	suite.Contains(resp.Results, suite.Response.Results[0])
	suite.Contains(resp.Results, suite.Response.Results[1])
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
