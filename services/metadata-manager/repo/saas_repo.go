package repo

import (
	"context"
	"database/sql"
	"fmt"

	kitlog "github.com/go-kit/kit/log"
	"local-testing.com/nk915/models"
)

type repository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func New(db *sql.DB, logger kitlog.Logger) error {
	return nil
	//	return &repository{
	//		db:     db,
	//		logger: kitlog.With(logger, "rep", "psqldb"),
	//	}, nil
}

func (repo *repository) CreateSaas(ctx context.Context, saas models.SaaS) error {
	// Run a transaction to sync the query model.
	//	err := crdb.ExecuteTx(ctx, repo.db, nil, func(tx *sql.Tx) error {
	//		return createOrder(tx, order)
	//	})
	//	if err != nil {
	//		return err
	//	}
	return nil
}

// 실제 쿼리 수행이 이루어지는 함수 : insert (삽입)
func createSaas() error {
	fmt.Println("createSaas")
	return nil
}

func (repo *repository) UpdateSaas(ctx context.Context, saas models.SaaS) error {
	fmt.Println("UpdateSaas")
	return nil
}

// GetSaasByID는 해당 ID를 통해 saas조회하는 쿼리를 수행합니다.
func (repo *repository) GetSaasByID(ctx context.Context, id string) error {
	fmt.Println("GetSaasByID")
	return nil
}
