package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/common/model"
)

func NewPsqlDB(cfg *config.Config, migrate bool) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.Postgres.PostgresqlHost,
		cfg.Postgres.PostgresqlUser,
		cfg.Postgres.PostgresqlPassword,
		cfg.Postgres.PostgresqlDbname,
		cfg.Postgres.PostgresqlPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if migrate {
		db.AutoMigrate(&model.User{})
		// if err != nil {
		// 	panic("Error when run migrations")
		// }
	}
	return db, nil
}
