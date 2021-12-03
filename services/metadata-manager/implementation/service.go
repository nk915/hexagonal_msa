package implementation

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"

	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/log/level"

	"local-testing.com/nk915/models"
)

var ErrEmpty = errors.New("empty string")

// service implements the SaaS Service
type service struct {
	//repository repo.Repository
	logger kitlog.Logger
}

// TODO: repo add
func NewService(logger kitlog.Logger) *service {
	return &service{
		logger: logger,
	}
}

//func NewService(rep repo.Repository, logger kitlog.Logger) Service {
//	return &service{
//		repository: rep,
//		logger:     logger,
//	}
//}

// Create makes on SaaS
func (s *service) Create(ctx context.Context, saas models.SaaS) (string, error) {
	logger := kitlog.With(s.logger, "method", "Create")
	uuid, _ := uuid.NewV4()
	id := uuid.String()

	// TODO: 임시 로그
	// TODO: CreateSaas 함수 호출 필요
	kitlevel.Info(logger).Log("info", "Call by Create : SaaS")

	return id, nil
}

// GetByID returns on saas given by id
func (s *service) GetByID(ctx context.Context, id string) (models.SaaS, error) {
	logger := kitlog.With(s.logger, "method", "GetByID")

	// TODO: CreateSaas 함수 호출 필요
	kitlevel.Info(logger).Log("Debug", "Call by GetByID : ", id)

	return models.SaaS{ID: "kng", Status: "on"}, nil
}

// Update makes on SaaS
//func (s *service) Update(ctx context.Context, id string) error {
//	logger := kitlog.With(s.logger, "method", "Update")
//
//	// TODO: UpdateSaas 함수 호출 필요
//	kitlevel.Info(logger).Log("Debug", "Call by Update : ", id)
//	return nil
//}
