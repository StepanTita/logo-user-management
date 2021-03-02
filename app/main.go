package app

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/logo-user-management/app/cors"
	"github.com/logo-user-management/app/ctx"
	"github.com/logo-user-management/app/data/postgres"
	"github.com/logo-user-management/app/logging"
	"github.com/logo-user-management/app/web"
	"github.com/logo-user-management/app/web/handlers"
	"github.com/logo-user-management/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

type App interface {
	Run() error
}

type app struct {
	log    *logrus.Logger
	config config.Config
}

func New(cfg config.Config) App {
	return &app{
		config: cfg,
		log:    cfg.Logging(),
	}
}

func (a *app) Run() error {
	defer func() {
		if rvr := recover(); rvr != nil {
			a.log.Error("app panicked\n", rvr)
		}
	}()

	a.log.WithField("port", a.config.Listener()).Info("Starting server")
	if err := http.ListenAndServe(a.config.Listener(), a.router()); err != nil {
		return errors.Wrap(err, "listener failed")
	}
	return nil
}

func (a *app) router() chi.Router {
	router := chi.NewRouter()

	router.Use(
		cors.Middleware(),
		logging.Middleware(a.log),
		ctx.Middleware(
			ctx.CtxLog(a.log),
			ctx.CtxConfig(a.config),
			ctx.CtxUsers(postgres.NewUsers(a.config)),
		),
	)

	router.Route("/logo/users", func(r chi.Router) {
		r.Route(fmt.Sprintf("/{%s}", web.UsernameRequestKey), func(r chi.Router) {
			r.Post("/auth", handlers.GetUser)
			r.Patch("/", handlers.UpdateUser)
			r.Delete("/", handlers.DeleteUser)
		})
		r.Post("/", handlers.CreateUser)
	})

	return router
}
