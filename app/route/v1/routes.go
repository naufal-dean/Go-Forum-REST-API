package v1

import (
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/controller/v1/auth"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/controller/v1/posts"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/controller/v1/threads"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/middleware"
)

func Setup(a *core.App) {
	// Open routes
	{
		// Get subrouter
		v1OpenR := a.Router.PathPrefix("/api/v1").Subrouter()

		// Auth routes
		v1OpenR.Handle("/register", auth.Register(a)).Methods("POST")
		v1OpenR.Handle("/login", auth.Login(a)).Methods("POST")
	}

	// Authenticated routes
	{
		// Get subrouter and set middleware
		v1AuthR := a.Router.PathPrefix("/api/v1").Subrouter()
		v1AuthR.Use(middleware.Auth(a))

		// Auth routes
		v1AuthR.Handle("/logout", auth.Logout(a)).Methods("POST")
		v1AuthR.Handle("/profile", auth.Profile(a)).Methods("POST")

		// Thread resource routes
		threadsR := v1AuthR.PathPrefix("/threads").Subrouter()
		threadsR.Handle("", threads.Create(a)).Methods("POST")
		threadsR.Handle("", threads.GetAll(a)).Methods("GET")
		threadsR.Handle("/{id}", threads.GetOne(a)).Methods("GET")
		threadsR.Handle("/{id}", threads.Update(a)).Methods("PUT")
		threadsR.Handle("/{id}", threads.Delete(a)).Methods("DELETE")

		// Posts resource routes
		postsR := v1AuthR.PathPrefix("/posts").Subrouter()
		postsR.Handle("", posts.Create(a)).Methods("POST")
		postsR.Handle("", posts.GetAll(a)).Methods("GET")
		postsR.Handle("/{id}", posts.GetOne(a)).Methods("GET")
		postsR.Handle("/{id}", posts.Update(a)).Methods("PUT")
		postsR.Handle("/{id}", posts.Delete(a)).Methods("DELETE")
	}
}
