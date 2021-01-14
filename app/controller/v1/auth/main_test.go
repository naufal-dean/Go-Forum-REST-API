package auth

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/test"
	"os"
	"testing"
)

var at *core.App

func init() {
	err := godotenv.Load("../../../../.test.env")
	if err != nil {
		log.Fatal("failed to load environment variable")
	}
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
