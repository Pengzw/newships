package data

import (
	"newships/app/department/internal/conf"
	"newships/pkg/db"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDeptRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{
		db: db.NewGormDB(c.Database.Driver, c.Database.Source),
	}, cleanup, nil
}
