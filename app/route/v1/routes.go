package v1

import (
	"github.com/naufal-dean/onboarding-dean-local/app/controller/v1/auth"
	"github.com/naufal-dean/onboarding-dean-local/app/controller/v1/posts"
	"github.com/naufal-dean/onboarding-dean-local/app/controller/v1/threads"
	"github.com/naufal-dean/onboarding-dean-local/app/core"
	"github.com/naufal-dean/onboarding-dean-local/app/middleware"
)

func Setup(a *core.App) {
	// Open routes start
	v1OpenR := a.Router.PathPrefix("/api/v1").Subrouter()

	// Auth routes
	v1OpenR.Handle("/register", auth.Register(a)).Methods("POST")
	v1OpenR.Handle("/login", auth.Login(a)).Methods("POST")


	// Authenticated routes start
	v1AuthR := a.Router.PathPrefix("/api/v1").Subrouter()
	v1AuthR.Use(middleware.Auth(a))

	// Auth routes
	v1AuthR.Handle("/logout", auth.Logout(a)).Methods("POST")

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
