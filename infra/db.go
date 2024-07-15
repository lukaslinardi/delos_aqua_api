package infra

import (
	"context"
	"database/sql"
	"time"

    _ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"

	constants "github.com/lukaslinardi/delos_aqua_api/domain/constants/general"
	log "github.com/sirupsen/logrus"
)

type Database interface {
	ConnectDB()
	Close()

	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	// DriverName() string

	Begin() (*sql.Tx, error)
	In(query string, params ...interface{}) (string, []interface{}, error)
	Rebind(query string) string
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row
	// QueryRowSqlx(query string, args ...interface{}) *sqlx.Row
	// QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	// GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type DatabaseList struct {
	Backend DatabaseType
}

type DatabaseType struct {
	Read  Database
	Write Database
}

type DBHandler struct {
	DB  *sqlx.DB
	Err error
	log *log.Logger
}

func NewDB(log *log.Logger) DBHandler {
	return DBHandler{
		log: log,
	}
}

func (d *DBHandler) ConnectDB() {

	dbs, err := sqlx.Open("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Error(constants.ConnectDBFail + " | " + err.Error())
		d.Err = err
		return
	}

	d.DB = dbs

	err = d.DB.Ping()
	if err != nil {
		log.Error(constants.ConnectDBFail, err.Error())
		d.Err = err
	}

	d.log.Info(constants.ConnectDBSuccess)
	d.DB.SetConnMaxLifetime(time.Duration(3600))
}

// Close - function for connection lost.
func (d *DBHandler) Close() {
	if err := d.DB.Close(); err != nil {
		d.log.Println(constants.ClosingDBFailed + " | " + err.Error())
	} else {
		d.log.Println(constants.ClosingDBSuccess)
	}
}

func (d *DBHandler) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := d.DB.Exec(query, args...)
	return result, err
}

func (d *DBHandler) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	result, err := d.DB.ExecContext(ctx, query, args...)
	return result, err
}

func (d *DBHandler) Query(query string, args ...interface{}) (*sql.Rows, error) {
	result, err := d.DB.Query(query, args...)
	return result, err
}

func (d *DBHandler) Select(dest interface{}, query string, args ...interface{}) error {
	err := d.DB.Select(dest, query, args...)
	return err
}

func (d *DBHandler) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	err := d.DB.SelectContext(ctx, dest, query, args...)
	return err
}

func (d *DBHandler) Get(dest interface{}, query string, args ...interface{}) error {
	err := d.DB.Get(dest, query, args...)
	return err
}

func (d *DBHandler) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	err := d.DB.GetContext(ctx, dest, query, args...)
	return err
}

func (d *DBHandler) Rebind(query string) string {
	return d.DB.Rebind(query)
}

func (d *DBHandler) In(query string, params ...interface{}) (string, []interface{}, error) {
	query, args, err := sqlx.In(query, params...)
	return query, args, err
}

func (d *DBHandler) Begin() (*sql.Tx, error) {
	return d.DB.Begin()
}

func (d *DBHandler) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return d.DB.QueryRowContext(ctx, query, args...)
}
