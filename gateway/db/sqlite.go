package db

import (
	"database/sql"

	"github.com/keshu12345/guardianlink/gateway/constant"
	_ "github.com/mattn/go-sqlite3"
	logger "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gopkg.in/gorp.v2"
)

var Module = fx.Options(
	fx.Provide(NewSQLiteInstnace),
)

type SQLite struct {
	fx.Out
	DB *gorp.DbMap
}

func NewSQLiteInstnace() (SQLite, error) {

	db, err := sql.Open(constant.Drivername.ToString(), constant.Databasename.ToString())
	if err != nil {
		logger.Errorf("Unable to open db file %#q", err)
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	// dbMap.AddTableWithName(model.Block{}, "blocks").SetKeys(false, "Height")
	// dbMap.AddTableWithName(model.User{}, "users").SetKeys(false, "Username")
	if err = dbMap.CreateTablesIfNotExists(); err != nil {
		return SQLite{}, err
	}
	return SQLite{DB: dbMap}, nil
}
