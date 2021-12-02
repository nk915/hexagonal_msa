package data

type ISaasService interface {
	GetSaasByID(string) (string, error)
}

type SaasService struct {
	ID, AccessKey, SecretKey string
}
