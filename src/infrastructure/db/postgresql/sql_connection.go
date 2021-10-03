package sql

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	maxIdleConns = 10
	maxOpenConns = 100
)

type Connection interface {
	GetConnection() *gorm.DB
}

type DBConnection struct {
	dialector gorm.Dialector
	options *DBOption
	connection *gorm.DB
}

func NewPostgreSQLConnection(opts ...*DBOption) *DBConnection {
	databaseOptions := MergeOptions(opts...)
	dialector := databaseOptions.getGormDialector()
	if dialector == nil {
		logrus.WithError(errors.New("empty connection url")).Fatalln("error creating connection")
		panic("Invalid dialector")
	}
	return &DBConnection{
		options:   databaseOptions,
		dialector: dialector,
	}
}

func (r *DBConnection) GetConnection() *gorm.DB {
	if r.connection == nil {
		newLogger := gormLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			gormLogger.Config{
				SlowThreshold: time.Second,
				LogLevel:      gormLogger.Warn,
				Colorful:      false,
			},
		)
		connection, err := gorm.Open(r.dialector, &gorm.Config{
			Logger: newLogger,
		})

		if err != nil {
			logrus.WithError(err).Errorln("error trying to connect to DB")
		} else {
			sqlDB, err := connection.DB()
			if err != nil {
				logrus.WithError(err).Errorln("error trying to connect to DB")
			}
			sqlDB.SetMaxIdleConns(maxIdleConns)
			sqlDB.SetMaxOpenConns(maxOpenConns)
			sqlDB.SetConnMaxLifetime(time.Hour)

			r.connection = connection
		}
	}
	return r.connection
}