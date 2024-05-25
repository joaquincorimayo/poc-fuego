package main

import (
	"github.com/go-fuego/fuego"
	"github.com/joaquincorimayo/poc-fuego/adapter"
	"github.com/joaquincorimayo/poc-fuego/api"
	"github.com/joaquincorimayo/poc-fuego/config"
)

func main() {
	conf, _ := config.NewConfig("./config/config.yaml")
	repo, _ := adapter.NewMongoRepository(
		conf.Database.URL,
		conf.Database.DB,
		conf.Database.Timeout)
	service := adapter.NewReportService(repo)
	handler := api.NewReportHandler(service)

	server := fuego.NewServer(
		fuego.WithAddr(conf.Server.Host),
		fuego.WithPort(conf.Server.Port),
		fuego.WithOpenAPIConfig(fuego.OpenAPIConfig{
			DisableSwagger: false,
		}),
	)

	fuego.Get(server, "/reports/{code}", handler.Get)
	fuego.Post(server, "/reports", handler.Post)

	server.Run()
}
