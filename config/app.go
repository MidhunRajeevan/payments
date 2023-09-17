package config

import (
	"os"
	"strconv"
)

type appConfig struct {
	ListenPort int
	SecretKey  string
}

var App appConfig

const (
	appListenPort = "APP_LISTEN_PORT"
	appSecretKey  = "APP_SECRET_KEY"
)

const (
	defaultListenPort = 9090
)

// InitializeApp Configuration
func InitializeApp() {
	var err error
	var ok bool

	// Listen Port
	if it, ok := os.LookupEnv(appListenPort); ok {
		if App.ListenPort, err = strconv.Atoi(it); err != nil {
			App.ListenPort = defaultListenPort
		}
	} else {
		App.ListenPort = defaultListenPort
	}

	// Secret Key
	if App.SecretKey, ok = os.LookupEnv(appSecretKey); !ok {
		panic("Please configure the secret key for encryption.")
	}
}
