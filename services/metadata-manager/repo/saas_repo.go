package repo

import (
	"context"
	"database/sql"
	"fmt"

	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
	"local-testing.com/nk915/models"
)

type repository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func New(db *sql.DB, logger kitlog.Logger) (Repository, error) {
	return &repository{
		db:     db,
		logger: kitlog.With(logger, "rep", "psqldb"),
	}, nil
}

func (repo *repository) InitTable() error {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS saas (id TEXT, status TEXT);`,
		//`CREATE TABLE user_addresses (address_id INT, user_id INT);`,
	}

	for _, table := range tables {
		_, err := repo.db.Exec(table)
		if err != nil {
			kitlevel.Error(repo.logger).Log("err", err.Error())
			return err
		}
		kitlevel.Info(repo.logger).Log("info", table)
	}
	return nil
}

func (repo *repository) CreateSaas(ctx context.Context, saas models.SaaS) error {
	query := `INSERT INTO saas (id, status) VALUES ($1, $2);`

	result, err := repo.db.Exec(query, saas.ID, saas.Status)
	if err != nil {
		kitlevel.Error(repo.logger).Log("err", err.Error())
		return err
	}
	kitlevel.Info(repo.logger).Log("info", result)
	return nil
}

func (repo *repository) UpdateSaas(ctx context.Context, saas models.SaaS) error {
	fmt.Println("UpdateSaas")
	return nil
}

// GetSaasByID는 해당 ID를 통해 saas조회하는 쿼리를 수행합니다.
func (repo *repository) GetSaasByID(ctx context.Context, id string) (models.SaaS, error) {
	var saasRow = models.SaaS{}
	query := `SELECT id, status FROM saas WHERE id = $1`

	if err := repo.db.QueryRow(query, id).
		Scan(&saasRow.ID, &saasRow.Status); err != nil {
		kitlevel.Error(repo.logger).Log("err", err.Error())
		return saasRow, err
	}

	kitlevel.Info(repo.logger).Log("info", saasRow)
	return saasRow, nil
}
