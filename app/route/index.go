package route

import (
	"github.com/naufal-dean/onboarding-dean-local/app/controller"
	"github.com/naufal-dean/onboarding-dean-local/app/core"
	"github.com/naufal-dean/onboarding-dean-local/app/middleware"
	v1 "github.com/naufal-dean/onboarding-dean-local/app/route/v1"
)

func Setup(a *core.App) {
	a.Router.Use(middleware.ErrorHandler)

	a.Router.Handle("/hello", controller.Hello(a))
	v1.Setup(a)
}
