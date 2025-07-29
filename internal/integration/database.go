package integration

import (
	"context"
	"leobelini/cashly/config"
	"leobelini/cashly/internal/types/database"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var Db *gorm.DB
var DbCtx context.Context

func StartDatabase() error {
	config.LoadDatabaseEnv()

	env := config.GetDatabaseEnv()

	var err error
	Db, err = gorm.Open(sqlite.Dialector{
		DriverName: "sqlite", // ou outro nome registrado
		DSN:        env.Filename,
	}, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}

	if env.AutoMigrate {
		Db.AutoMigrate(&database.User{})
	}
	DbCtx = context.Background()

	return nil
}
