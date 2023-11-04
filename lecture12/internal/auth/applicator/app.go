package applicator

import (
	"context"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/config"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/repository"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/server/http"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/service"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/transport"
	"log"
	"os"
	"os/signal"
)

type App struct {
	cfg *config.Config
}

func New(cfg *config.Config) *App {
	return &App{cfg: cfg}
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	gracefullyShutdown(cancel)
	repo, err := repository.NewRepository(ctx, a.cfg)
	if err != nil {
		log.Fatalf("cannot сonnect to mainDB '%s:%d': %v", a.cfg.Database.Host, a.cfg.Database.Port, err)
	}
	userTransport := transport.NewTransport(a.cfg.Transport.User)
	authService := service.NewManager(repo, a.cfg.Auth, userTransport)
	endPointHandler := http.NewManager(authService)
	HTTPServer := http.NewServer(a.cfg, endPointHandler)
	return HTTPServer.StartHTTPServer(ctx)
}
func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}
