package config

import (
	"os"

	"github.com/TwiN/logr"
)

func loadEnvVarSettings(conf *Config) {
	frontendUrl, ok := os.LookupEnv("FRONTEND_URL")
	if ok {
		(*conf).Urls.FrontendURL = frontendUrl
	}
	backendUrl, ok := os.LookupEnv("BACKEND_URL")
	if ok {
		(*conf).Urls.BackendURL = backendUrl
	}
	defaultRedirect, ok := os.LookupEnv("DEFAULT_REDIRECT")
	if ok {
		(*conf).Urls.DefaultRedirect = defaultRedirect
	}

	testingMode, ok := os.LookupEnv("TEST_MODE")
	if ok && testingMode == "true" {
		logr.Warn("Testing mode enabled")
		(*conf).Urls.DefaultRedirect = "/"
		(*conf).TestingMode = true
	}
}
