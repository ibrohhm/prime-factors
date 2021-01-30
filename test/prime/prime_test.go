package prime_test

import (
	"testing"

	"github.com/prime-factors/prime"
	"github.com/stretchr/testify/suite"
)

type PrimeTestSuite struct {
	suite.Suite
	prime prime.Prime
}

func (suite *PrimeTestSuite) SetupTest() {
	suite.prime = prime.Prime{}
}

func (suite *PrimeTestSuite) TestPrimeCheck() {
	suite.Equal(suite.prime.Check(5), true)   // prime number
	suite.Equal(suite.prime.Check(12), false) // composite number
}

func (suite *PrimeTestSuite) TestPrimeFactor() {
	suite.Equal(suite.prime.Factors(5), []int{5})     // sprime number
	suite.Equal(suite.prime.Factors(12), []int{2, 3}) // composite number
}

func TestPrimeTestSuite(t *testing.T) {
	suite.Run(t, new(PrimeTestSuite))
}
