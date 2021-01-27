package auth

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/naufal-dean/go-forum-rest-api/app/core"
	"github.com/naufal-dean/go-forum-rest-api/app/test"
	"os"
	"testing"
)

var at *core.App

func init() {
	_ = godotenv.Load("../../../../.test.env")
	a, err := test.NewTestApp()
	if err != nil {
		log.Fatal(err)
	}
	at = a
}

func setup() {
	err := test.DatabaseUp(at)
	if err != nil {
		log.Fatal(err)
	}
}

func teardown() {
	test.DatabaseDown(at)
}

func TestMain(m *testing.M) {
	code := m.Run()
	teardown() // assert table is dropped
	os.Exit(code)
}
