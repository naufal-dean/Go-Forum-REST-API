package route

import (
	"github.com/naufal-dean/onboarding-dean-local/app/controller"
	"github.com/naufal-dean/onboarding-dean-local/app/core"
)

func Setup(a *core.App) {
	a.Router.Handle("/", controller.Hello(a))
}
