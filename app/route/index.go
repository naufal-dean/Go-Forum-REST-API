package route

import (
	"github.com/pkg/errors"
	"github.com/rakyll/statik/fs"
	"github.com/naufal-dean/go-forum-rest-api/app/controller"
	"github.com/naufal-dean/go-forum-rest-api/app/core"
	"github.com/naufal-dean/go-forum-rest-api/app/middleware"
	v1 "github.com/naufal-dean/go-forum-rest-api/app/route/v1"
	_ "github.com/naufal-dean/go-forum-rest-api/app/static/swaggerui/statik"
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
		panic(errors.Wrap(err, "failed to create new statik file server instance"))
	}
	a.Router.PathPrefix("/api/docs/").Handler(http.StripPrefix("/api/docs/", http.FileServer(statikFS)))
}
