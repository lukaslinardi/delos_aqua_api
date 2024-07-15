package db

import (
	"github.com/lukaslinardi/delos_aqua_api/infra"
	"github.com/sirupsen/logrus"
)

type Database struct {
	Farm Farm
}

func NewDatabase(db *infra.DatabaseList, logger *logrus.Logger) Database {
	return Database{
		Farm: newFarm(db, logger),
	}
}
