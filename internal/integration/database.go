package integration

import (
	"leobelini/cashly/config"
	domainUser "leobelini/cashly/internal/domain/user"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func StartDatabase() (*gorm.DB, error) {
	config.LoadDatabaseEnv()

	env := config.GetDatabaseEnv()

	db, err := gorm.Open(sqlite.Dialector{
		DriverName: "sqlite", // ou outro nome registrado
		DSN:        env.Filename,
	}, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	if env.AutoMigrate {
		db.AutoMigrate(&domainUser.User{})
	}

	return db, nil
}
