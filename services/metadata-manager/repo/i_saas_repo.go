package repo

import (
	"context"

	"local-testing.com/nk915/models"
)

type Repository interface {
	InitTable() error
	CreateSaas(ctx context.Context, saas models.SaaS) error
	UpdateSaas(ctx context.Context, saas models.SaaS) error
	GetSaasByID(ctx context.Context, id string) (models.SaaS, error)
}
