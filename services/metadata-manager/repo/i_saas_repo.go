package repo

import (
	"context"

	"local-testing.com/nk915/models"
)

type Repository interface {
	CreateSaas(ctx context.Context, saas models.SaaS)
	UpdateSaas(ctx context.Context, saas models.SaaS)
	GetSaasByID(ctx context.Context, id string)
}
