package database

import (
	"context"
	"database/sql"
)

type IConnection interface {
	Exec(c context.Context, query string, sqlParams ...interface{}) (sql.Result, error)
	QueryRow(c context.Context, query string, sqlParams ...interface{}) *sql.Row
}

type Connection struct {
	DBConnection *sql.DB
}

var DBConnection IConnection

func NewConnection(dbConnectionObj *sql.DB) {
	DBConnection = &Connection{
		DBConnection: dbConnectionObj,
	}
}

func (dbConnection *Connection) Exec(c context.Context, query string, sqlParams ...interface{}) (sql.Result, error) {
	dbResult, err := dbConnection.DBConnection.Exec(query, sqlParams...)

	if err != nil {
		return nil, err
	}

	return dbResult, nil
}

func (dbConnection *Connection) QueryRow(c context.Context, query string, sqlParams ...interface{}) *sql.Row {
	dbRow := dbConnection.DBConnection.QueryRow(query, sqlParams...)
	return dbRow
}
