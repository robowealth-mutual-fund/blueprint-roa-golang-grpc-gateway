package test

import "errors"

func (suite *PackageTestSuite) TestCreate() {
	givenInput := suite.makeTestCreateInput()
	resultRepoCart := suite.makeTestCart()
	suite.repo.On("Create", resultRepoCart).Once().Return(nil)
	name, err := suite.service.Create(suite.ctx, givenInput)
	suite.Equal("aaa", name)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestCreateError() {
	givenInput := suite.makeTestCreateInput()
	resultRepoCart := suite.makeTestCart()
	suite.repo.On("Create", resultRepoCart).Once().Return(errors.New("error"))
	_, err := suite.service.Create(suite.ctx, givenInput)
	suite.Error(errors.New("error"), err)
}
