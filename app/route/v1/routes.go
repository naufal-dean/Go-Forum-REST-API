package v1

import (
	"github.com/naufal-dean/onboarding-dean-local/app/controller/v1/posts"
	"github.com/naufal-dean/onboarding-dean-local/app/controller/v1/threads"
	"github.com/naufal-dean/onboarding-dean-local/app/core"
)

func Setup(a *core.App) {
	// Get subrouter
	v1 := a.Router.PathPrefix("/api/v1").Subrouter()

	// Posts resource routes
	postsR := v1.PathPrefix("/posts").Subrouter()
	postsR.Handle("", posts.Create(a)).Methods("POST")
	postsR.Handle("", posts.GetAll(a)).Methods("GET")
	postsR.Handle("/{id}", posts.GetOne(a)).Methods("GET")
	postsR.Handle("/{id}", posts.Update(a)).Methods("PUT")
	postsR.Handle("/{id}", posts.Delete(a)).Methods("DELETE")

	// Thread resource routes
	threadsR := v1.PathPrefix("/threads").Subrouter()
	threadsR.Handle("", threads.Create(a)).Methods("POST")
	threadsR.Handle("", threads.GetAll(a)).Methods("GET")
	threadsR.Handle("/{id}", threads.GetOne(a)).Methods("GET")
	threadsR.Handle("/{id}", threads.Update(a)).Methods("PUT")
	threadsR.Handle("/{id}", threads.Delete(a)).Methods("DELETE")
}
