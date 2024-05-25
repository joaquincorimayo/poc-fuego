package usecase

import (
	"github.com/joaquincorimayo/poc-fuego/domain/model"
)

// Service interface to all CRUD operations
type Service interface {
	FindByCode(code string) (*model.Report, error)
	Save(report *model.Report) error
}
