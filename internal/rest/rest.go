package rest

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	api "github.com/jorgesanchez-e/test/internal/service"
)

const (
	messagesPath string = "/messages"
)

type Server struct {
	e *echo.Echo
}

func NewServer(ctx context.Context) Server {
	return Server{
		e: echo.New(),
	}
}

func (s Server) Start() {
	s.initApi()

	err := s.e.Start(":8080")
	if err != nil {
		logrus.Error(err)
	}
}

func (s Server) initApi() {
	s.e.GET(messagesPath, service)
}

func service(c echo.Context) error {
	ctx := context.Background()

	s, err := api.GetResults(ctx)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, s, " ")
}
