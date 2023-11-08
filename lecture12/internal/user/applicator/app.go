package applicator

import (
	"context"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/config"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/repository"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/server/http"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/server/http/middleware"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/service"
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
	authService, _ := service.NewManager(repo)
	endPointHandler := http.NewManager(authService)
	jwt := middleware.NewJWTAuth(a.cfg)
	HTTPServer := http.NewServer(a.cfg, endPointHandler, jwt)
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
