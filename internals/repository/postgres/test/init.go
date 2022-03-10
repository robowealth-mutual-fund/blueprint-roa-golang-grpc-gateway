//go:build integration
// +build integration

package test

import (
	"context"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/entity"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/infrastructure/database"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/repository/postgres"
	"github.com/stretchr/testify/suite"
)

type Filter struct {
	PlatformID string `gorm:"type:varchar(255)" json:"platform_id"`
	Subject    string `gorm:"type:varchar(255)" json:"subject"`
	Metadata   string `gorm:"type:JSONB" json:"metadata"`
}

type PackageTestSuite struct {
	suite.Suite
	ctx     context.Context
	repo    *postgres.PostgresRepository
	verbose *entity.Product
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	conf := config.NewConfiguration()
	connetBase := database.NewServerBase(conf)
	suite.repo = postgres.NewRepository(connetBase)
}

func (suite *PackageTestSuite) SetupTest() {
	conf := config.NewConfiguration()
	connetBase := database.NewServerBase(conf)
	suite.repo = postgres.NewRepository(connetBase)
}

func (suite *PackageTestSuite) makeTestStruct(name string, detail string) (test *entity.Product) {
	return &entity.Product{
		Name:   name,
		Detail: detail,
	}
}
