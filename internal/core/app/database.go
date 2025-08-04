package app

import (
	"leobelini/cashly/internal/core/dto"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

type Database struct {
	env *dto.DtoEnvApp
	Db  *gorm.DB
}

func NewDatabase(env *dto.DtoEnvApp) *Database {
	return &Database{env: env}
}

func (c *Database) Start() error {

	db, err := gorm.Open(sqlite.Dialector{
		DriverName: "sqlite",
		DSN:        c.env.Database.File,
	}, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}

	c.Db = db

	return nil
}

func (s *Database) Close() {
	if s.Db != nil {
		db, _ := s.Db.DB()
		db.Close()
	}
}

func (s *Database) Migrate(dst ...interface{}) {
	if s.Db != nil {
		s.Db.AutoMigrate(dst...)
	}
}
