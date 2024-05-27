package cart_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/cart/test"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
