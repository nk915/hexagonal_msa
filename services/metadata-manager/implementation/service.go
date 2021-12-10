package implementation

import (
	"context"
	"database/sql"
	"errors"

	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/log/level"

	"local-testing.com/nk915/models"
	svc "local-testing.com/nk915/repo"
)

var ErrEmpty = errors.New("empty string")

// service implements the SaaS Service
type service struct {
	repository svc.Repository
	logger     kitlog.Logger
}

func NewService(rep svc.Repository, logger kitlog.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

// Create makes on SaaS
func (s *service) Create(ctx context.Context, saas models.SaaS) (string, error) {
	logger := kitlog.With(s.logger, "method", "Create")
	//	uuid, _ := uuid.NewV4()
	//	id := uuid.String()

	if err := s.repository.CreateSaas(ctx, saas); err != nil {
		kitlevel.Error(logger).Log("err", err)
		return "", ErrCmdRepository
	}

	// TODO: 추후 제거
	kitlevel.Info(logger).Log("info", "Call by Create : SaaS")
	kitlevel.Info(logger).Log("info", "Create (SaaS ID) : "+saas.ID)
	kitlevel.Info(logger).Log("info", "Create (SaaS Status) : "+saas.Status)
	return saas.ID, nil
}

// GetByID returns on saas given by id
func (s *service) GetByID(ctx context.Context, id string) (models.SaaS, error) {
	logger := kitlog.With(s.logger, "method", "GetByID")

	// TODO: 추후 제거
	kitlevel.Info(logger).Log("Debug", "Call by GetByID : ", id)

	saas, err := s.repository.GetSaasByID(ctx, id)
	if err != nil {
		kitlevel.Error(logger).Log("err", err)
		if err == sql.ErrNoRows {
			return saas, ErrSaasNotFound
		}
		return saas, ErrQueryRepository
	}

	return saas, nil
	// return models.SaaS{ID: "kng", Status: "on"}, nil
}

// Update makes on SaaS
//func (s *service) Update(ctx context.Context, id string) error {
//	logger := kitlog.With(s.logger, "method", "Update")
//
//	// TODO: UpdateSaas 함수 호출 필요
//	kitlevel.Info(logger).Log("Debug", "Call by Update : ", id)
//	return nil
//}
