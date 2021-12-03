package endpoint

import (
	"local-testing.com/nk915/models"
)

type CreateRequest struct {
	SaaS models.SaaS
}

type CreateResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error"`
}

type GetByIDRequest struct {
	ID string
}

type GetByIDResponse struct {
	SaaS models.SaaS `json:"saas"`
	Err  error       `json:"error"`
}
