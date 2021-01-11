package route

import (
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/controller"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/middleware"
	v1 "gitlab.com/pinvest/internships/hydra/onboarding-dean/app/route/v1"
)

func Setup(a *core.App) {
	a.Router.Use(middleware.Logger)
	a.Router.Use(middleware.ErrorHandler)

	a.Router.Handle("/hello", controller.Hello(a))
	v1.Setup(a)
}
