package applicator

import (
	"context"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/config"
)

type App struct {
	cfg *config.Config
}

func New(cfg *config.Config) *App {
	return &App{cfg: cfg}
}

func (a *App) Run(ctx context.Context) {

}
