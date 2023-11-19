package http

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/config"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/server/http/middleware"
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"time"
)

type Server struct {
	cfg     *config.Config
	handler *Manager
	App     *echo.Echo
	m       *middleware.JWTAuth
}

func NewServer(cfg *config.Config, handler *Manager, jwt *middleware.JWTAuth) *Server {

	return &Server{
		cfg:     cfg,
		handler: handler,
		m:       jwt,
	}
}

func (s *Server) StartHTTPServer(ctx context.Context) error {
	s.App = s.BuildEngine()
	s.SetupRoutes()
	go func() {
		if err := s.App.Start(fmt.Sprintf(":%v", s.cfg.HttpServer.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%v\n", err)
		}
	}()
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := s.App.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%v", err)
	}
	log.Print("server exited properly")
	return nil
}

func (s *Server) BuildEngine() *echo.Echo {
	e := echo.New()
	e.Use(middleware2.CORSWithConfig(middleware2.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	return e
}
