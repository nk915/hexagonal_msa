package implementation

import (
	"context"
	"errors"

	"local-testing.com/nk915/models"
)

var (
	ErrSaasNotFound    = errors.New("SaaS not found")
	ErrCmdRepository   = errors.New("unable to command repository")
	ErrQueryRepository = errors.New("unable to query repository")
)

type Service interface {
	Create(ctx context.Context, saas models.SaaS) (string, error)
	GetByID(ctx context.Context, id string) (models.SaaS, error)
	//	Update(ctx context.Context, saas models.SaaS) error
}
