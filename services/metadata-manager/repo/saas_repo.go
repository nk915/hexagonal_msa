package repo

import (
	"database/sql"

	kitlog "github.com/go-kit/kit/log"
)

type repository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func New(db *sql.DB, logger kitlog.Logger) (Repository, error) {

	return &repository, nil
	//	return &repository{
	//		db:     db,
	//		logger: kitlog.With(logger, "rep", "psqldb"),
	//	}, nil
}
