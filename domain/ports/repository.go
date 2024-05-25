package ports

import (
	"github.com/joaquincorimayo/poc-fuego/domain/model"
)

// Repository interface port to database
type Repository interface {
	FindByCode(code string) (*model.Report, error)
	Save(report *model.Report) error
}
