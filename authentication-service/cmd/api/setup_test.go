package main

import (
	"calvarado2004/microservices-go/authentication/data"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {

	repo := data.NewPostgresTestRepository(nil)
	testApp.Repo = repo

	code := m.Run()
	os.Exit(code)
}
