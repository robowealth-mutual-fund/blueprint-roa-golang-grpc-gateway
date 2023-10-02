//go:build integration
// +build integration

package postgres_test

import (
	"testing"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/repository/postgres/test"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
