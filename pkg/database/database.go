package database

import (
	"database/sql"
	"fmt"
	"github.com/ahmadrezamusthafa/assessment/common/logger"
	"github.com/ahmadrezamusthafa/assessment/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Result struct {
	Data  interface{}
	Error error
}

type Block func(db *sqlx.Tx, c chan Result)

type Database struct {
	Config config.Config `inject:"config"`
	db     *sqlx.DB
}

func (m Database) GetDB() *sqlx.DB {
	return m.db
}

func (m Database) WithTransaction(block Block) (result Result, err error) {
	c := make(chan Result)
	tx, err := m.db.Beginx()
	if err == nil {
		go block(tx, c)
	} else {
		logger.Err(err.Error())
		return
	}
	result = <-c
	if result.Error != nil {
		tx.Tx.Rollback()
	} else {
		tx.Tx.Commit()
	}
	return
}

func (m Database) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	return m.db.BindNamed(query, arg)
}

func (m Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return m.db.QueryRow(query, args)
}

func (m Database) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return m.db.NamedExec(query, arg)
}

func (m Database) Get(dest interface{}, query string, args ...interface{}) (err error) {
	return m.db.Get(dest, query, args...)
}

func (m Database) Select(dest interface{}, query string, args ...interface{}) (err error) {
	return m.db.Select(dest, query, args...)
}

func (m Database) In(query string, params map[string]interface{}) (string, []interface{}, error) {
	query, args, err := sqlx.Named(query, params)
	if err != nil {
		return query, args, err
	}
	return sqlx.In(query, args...)
}

func (m Database) Prepare(query string) (*sqlx.NamedStmt, error) {
	return m.db.PrepareNamed(query)
}

func (m Database) PrepareBind(query string) (*sqlx.Stmt, error) {
	return m.db.Preparex(query)
}

func (m Database) Rebind(query string) string {
	return m.db.Rebind(query)
}

func (m Database) IsReady() bool {
	if m.db == nil {
		return false
	}
	if err := m.db.Ping(); err != nil {
		logger.Err(err.Error())
		return false
	}
	return true
}

type AssessmentDatabase struct {
	Config config.Config `inject:"config"`
	Database
}

func (m *AssessmentDatabase) Shutdown() {
	if m.db != nil {
		logger.Info("Closing database connection...")
		m.db.Close()
	}
}

func (m *AssessmentDatabase) StartUp() {
	logger.Info("Init database connection...")
	conf := m.Config
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		conf.DatabaseHost, conf.DatabasePort, conf.DatabaseName, conf.DatabaseUsername, conf.DatabasePassword)
	info := fmt.Sprintf("%s:%s db=%s", conf.DatabaseHost, conf.DatabasePort, conf.DatabaseName)

	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		logger.Warn("Failed to connect [%s]", info)
	} else if err := db.Ping(); err != nil {
		logger.Err("Error while connecting to [%s]", info)
	} else {
		logger.Info("Successfully connected to postgres [%s]", info)
	}
	m.db = db
}
