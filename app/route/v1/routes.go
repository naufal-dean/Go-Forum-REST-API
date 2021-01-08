package v1

import (
	"github.com/naufal-dean/onboarding-dean-local/app/controller/v1/posts"
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
}

