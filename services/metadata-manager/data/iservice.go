package data

type ISaasService interface {
	GetSaasInfo(string) (string, error)
}

type SaasService struct {
	//ID, AccessKey, SecretKey string
}
