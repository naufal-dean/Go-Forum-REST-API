package app

import (
	"fmt"
	"github.com/naufal-dean/go-forum-rest-api/app/core"
	"net/http"
)

func Exec(flags map[string]bool) {
	var a = &core.App{}
	InitApp(a, flags)

	addr := ":8080"
	fmt.Printf("[+] Server started at %s...\n", addr)
	http.ListenAndServe(addr, a.Router)
}

