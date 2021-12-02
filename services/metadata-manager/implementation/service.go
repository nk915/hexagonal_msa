package implementation

import (
	"errors"

	"local-testing.com/nk915/data"
)

var ErrEmpty = errors.New("empty string")

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *data.SaasService) GetSaasInfo(_key string) (string, error) {
	if _key == "" {
		return "", ErrEmpty
	}
	return "Hello World", nil
}
