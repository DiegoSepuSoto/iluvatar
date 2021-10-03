package sql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBOption struct {
	sqlDialect   SQLDialect
	databaseName *string
	host         *string
	port         *int
	user         *string
	password     *string
}

func Config() *DBOption {
	return new(DBOption)
}

func (c *DBOption) SetSQLDialect(sqlDialect SQLDialect) *DBOption {
	c.sqlDialect = sqlDialect
	return c
}

func (c *DBOption) DatabaseName(databaseName string) *DBOption {
	c.databaseName = &databaseName
	return c
}

func (c *DBOption) Host(host string) *DBOption {
	c.host = &host
	return c
}

func (c *DBOption) Port(port int) *DBOption {
	c.port = &port
	return c
}

func (c *DBOption) User(user string) *DBOption {
	c.user = &user
	return c
}

func (c *DBOption) Password(password string) *DBOption {
	c.password = &password
	return c
}

func MergeOptions(opts ...*DBOption) *DBOption {
	option := new(DBOption)

	for _, opt := range opts {
		if opt.sqlDialect != "" {
			option.sqlDialect = opt.sqlDialect
		} else {
			panic("Invalid sqlDialect")
		}
		if opt.databaseName != nil {
			option.databaseName = opt.databaseName
		}
		if opt.host != nil {
			option.host = opt.host
		}
		if opt.port != nil {
			option.port = opt.port
		}
		if opt.user != nil {
			option.user = opt.user
		}
		if opt.password != nil {
			option.password = opt.password
		}
	}
	return option
}

const (
	URLPostgreSQLServerFormat = "host=%s port=%d user=%s password=%s dbname=%s"

)

type SQLDialect string

const (
	PostgresQL SQLDialect = "postgres"
)

func (c *DBOption) getGormDialector() gorm.Dialector {
	switch c.sqlDialect {
	case PostgresQL:
		url := fmt.Sprintf(URLPostgreSQLServerFormat, *c.host, *c.port, *c.user, *c.password, *c.databaseName)
		return postgres.Open(url)
	default:
		return nil
	}
}
