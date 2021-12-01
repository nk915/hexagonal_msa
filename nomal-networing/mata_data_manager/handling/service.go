package handling

import (
	"errors"
)

type IsaasService interface {
	getSaasInfo(string) (string, error)
}

type saasService struct{}

func (saasService) getSaasInfo(_key string) (string, error) {
	if _key == "" {
		return "", ErrEmpty
	}
	return "Hello World", nil
}

var ErrEmpty = errors.New("empty string")
