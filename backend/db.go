package main

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DatabaseConfiguration struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Extras   string `yaml:"extras"`
}

func (dc *DatabaseConfiguration) SafeDsn() string {
	return fmt.Sprintf("postgres://%s:[REDACTED]@%s:%s/%s?%s", dc.Username, dc.Host, dc.Port, dc.Database, dc.Extras)
}

func NewDatabase(cfg DatabaseConfiguration) *bun.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Extras)
	//dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	return bun.NewDB(sqlDb, pgdialect.New())
}
