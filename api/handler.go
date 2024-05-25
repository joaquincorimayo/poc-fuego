package api

import (
	"github.com/go-fuego/fuego"
	"github.com/joaquincorimayo/poc-fuego/domain/model"
)

type ReportHandler interface {
	Get(c *fuego.ContextNoBody) (model.Report, error)
	Post(c *fuego.ContextWithBody[model.Report]) (model.ReportCreatedResponse, error)
}
