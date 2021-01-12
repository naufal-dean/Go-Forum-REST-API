package route

import (
	"github.com/rakyll/statik/fs"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/controller"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/middleware"
	v1 "gitlab.com/pinvest/internships/hydra/onboarding-dean/app/route/v1"
	_ "gitlab.com/pinvest/internships/hydra/onboarding-dean/app/static/swaggerui/statik"
	"net/http"
)

func Setup(a *core.App) {
	a.Router.Use(middleware.Logger)
	a.Router.Use(middleware.ErrorHandler)

	// Hello world
	a.Router.Handle("/hello", controller.Hello(a))

	// API
	v1.Setup(a)

	// API docs
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	a.Router.PathPrefix("/docs").Handler(http.StripPrefix("/docs", http.FileServer(statikFS)))
}
