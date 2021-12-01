package implementation

import (
	"errors"

	"github.com/nk915/k8s_msa_example/services/data"
)

var ErrEmpty = errors.New("empty string")

func (s *data.SaasService) GetSaasInfo(_key string) (string, error) {
	if _key == "" {
		return "", ErrEmpty
	}
	return "Hello World", nil
}
