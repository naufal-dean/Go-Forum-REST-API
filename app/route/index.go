package route

import (
	"github.com/naufal-dean/onboarding-dean-local/app/controller"
	"github.com/naufal-dean/onboarding-dean-local/app/core"
	v1 "github.com/naufal-dean/onboarding-dean-local/app/route/v1"
)

func Setup(a *core.App) {
	a.Router.Handle("/hello", controller.Hello(a))
	v1.Setup(a)
}
