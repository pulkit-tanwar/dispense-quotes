package api

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/pulkit-tanwar/dispense-quotes/lib/config"
	"github.com/pulkit-tanwar/dispense-quotes/lib/quoteslogic"
)

type Server struct {
	*config.Config
	router    *echo.Echo
	dispenser quoteslogic.Dispenser
}

func NewServer(cfg *config.Config, dispenser quoteslogic.Dispenser) *Server {
	server := &Server{
		Config:    cfg,
		router:    echo.New(),
		dispenser: dispenser,
	}

	server.router.GET(path.Join(cfg.APIPath, "/healthcheck"), server.Ping)
	server.router.GET(path.Join(cfg.APIPath, "/version"), server.Version)
	server.router.GET(path.Join(cfg.APIPath, "/quote"), server.Quote)

	server.router.Use(middleware.Logger())
	server.router.HideBanner = true
	server.router.HidePort = true

	return server
}

// Start - This function will start the HTTP server
func (s *Server) Start() error {
	address := fmt.Sprintf("%s:%d", s.Host, s.Port)
	log.Infof("Listening on %s", address)
	return s.router.Start(address)
}

// Start - This function will stop the HTTP server
func (s *Server) Stop(ctx context.Context) error {
	return s.router.Shutdown(ctx)
}

// Start - This function will ping the server
func (s *Server) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"service": "up"})
}

// Version - Returns the version
func (s *Server) Version(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"version": quoteslogic.Version})
}

// Quote - Return a random quote in JSON
func (s *Server) Quote(c echo.Context) error {
	return c.JSON(http.StatusOK, s.dispenser.RandomQuote())
}
