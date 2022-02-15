package database

import (
	"database/sql"
	"fmt"
	"net/url"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/ms-user-portal/app/config"
	"github.com/ms-user-portal/app/logging"
)

const connectionTimeout = 15

func Initialize(cfg *config.Configuration) {
	lw := logging.LogForFunc()
	if cfg == nil {
		lw.Error("Config for database is nil")
	}

	queryParams := url.Values{}
	queryParams.Add("connection timeout", strconv.Itoa(connectionTimeout))
	queryParams.Add("database", cfg.Database.DBName)
	queryParams.Add("app name", cfg.MSName)

	u := &url.URL{
		Scheme:   cfg.Database.Type,
		User:     url.UserPassword(cfg.Database.UserName, cfg.Database.Password),
		Host:     fmt.Sprintf("%s:%d", cfg.Database.Server, cfg.Database.Port),
		RawQuery: queryParams.Encode(),
	}

	dbConnection, err := sql.Open("sqlserver", u.String())

	if dbConnection == nil || err != nil {
		lw.WithField("error", err.Error()).Panic()
		return
	}
	lw.Info("Database opened successfully")

	err = dbConnection.Ping()
	if err != nil {
		lw.WithError(err).Fatal("Cannot connect to database")
		return
	}

	NewConnection(dbConnection)
	lw.Info("Database connection established successfully")
}
