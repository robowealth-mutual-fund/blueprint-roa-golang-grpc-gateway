//go:build integration
// +build integration

package test

func (suite *PackageTestSuite) TestCreate() {
	input := suite.makeTestStruct("odini", "odini01121")
	err := suite.repo.Create(input)
	suite.NoError(err)
}
