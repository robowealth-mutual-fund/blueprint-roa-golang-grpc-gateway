package container

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	controllerUsers "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/users"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/users"
	"github.com/robowealth-mutual-fund/shared-utility/validator"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller"
	controllerCart "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/cart"
	controllerCategory "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/category"
	controllerProduct "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/product"
	warehouseController "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/warehouse"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/infrastructure/database"
	grpcServer "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/infrastructure/grpcServer"
	httpServer "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/infrastructure/httpServer"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/infrastructure/jaeger"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/repository/postgres"
	cartService "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/cart"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/category"
	serviceProduct "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/product"
	warehouseService "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/warehouse"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/utils"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/utils/logrus"
)

type Container struct {
	container *dig.Container
}

func (c *Container) Configure() error {

	servicesConstructors := []interface{}{
		config.NewConfiguration,
		controllerCart.NewController,
		cartService.NewService,
		grpcServer.NewServer,
		database.NewServerBase,
		http.NewServeMux,
		httpServer.NewServer,
		runtime.NewServeMux,
		jaeger.NewJaeger,

		logrus.NewLog,
		controller.NewHealthZController,
		controller.NewPingPongController,
		validator.NewCustomValidator,
		controllerProduct.NewController,
		serviceProduct.NewService,
		postgres.NewRepository,
		utils.NewUtils,
		utils.NewCustomValidator,
		category.NewService,
		warehouseService.NewService,
		users.NewService,
		controllerUsers.NewController,
		controllerCategory.NewController,
		warehouseController.NewController,
	}
	for _, service := range servicesConstructors {
		if err := c.container.Provide(service); err != nil {
			return err
		}
	}
	appConfig := config.NewConfiguration()
	jaeger.NewJaeger(appConfig)
	logrus.NewLog()
	return nil
}

func (c *Container) Start() error {
	log.Info("Start Container")
	if err := c.container.Invoke(func(s *grpcServer.Server, h *httpServer.Server, v *validator.CustomValidator) {
		go func() {
			_ = h.Start()
		}()
		s.Start()

	}); err != nil {
		log.Errorf("%s", err)

		return err
	}

	return nil
}

// MigrateDB ...
func (c *Container) MigrateDB() error {
	log.Info("Start Container DB")

	if err := c.container.Invoke(func(d *database.DB) {
		d.MigrateDB()
	}); err != nil {
		return err
	}

	return nil
}

func NewContainer() (*Container, error) {
	d := dig.New()

	container := &Container{
		container: d,
	}

	if err := container.Configure(); err != nil {
		return nil, err
	}

	return container, nil
}
