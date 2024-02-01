package config

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"

	"git.motiolabs.com/library/motiolibs/merror"
	"github.com/spf13/viper"
)

type Config struct {
	AppMode      string
	AppName      string
	AppUrl       string
	AppClientUrl string
	AppApiKey    string
	PgHostname   string
	PgUsername   string
	PgPassword   string
	PgDatabase   string
	PgTimezone   string
	PgSslMode    string
	JwtKey       string
	JwtExpired   int
	JwtRefresh   int
	AppPort      int
	PgPort       int
	AppDebug     bool

	// ADDITIONAL
	SMTPServer     string
	SMTPPort       int
	SMTPUsername   string
	SMTPPassword   string
	SenderEmail    string
	RecipientEmail string
	SubjectEmail   string
}

func load() Config {
	//find file location
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to get the current filename")
	}
	p := filepath.Dir(filename)

	// Set the file type
	viper.SetConfigType("toml")
	// Set the file name of the configurations file
	viper.SetConfigName("config")
	// Set the path to look for the configurations file
	viper.AddConfigPath(filepath.Clean(filepath.Join(p, "..")))

	// Read in the configuration file
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	cfg := Config{
		AppMode:      viper.GetString("app.mode"),
		AppClientUrl: viper.GetString("app.client_url"),
		AppDebug:     viper.GetBool("app.debug"),
		AppName:      viper.GetString("app.name"),
		AppUrl:       viper.GetString("app.url"),
		AppPort:      viper.GetInt("app.port"),
		AppApiKey:    viper.GetString("app.api_key"),
		PgHostname:   viper.GetString("postgresql.hostname"),
		PgPort:       viper.GetInt("postgresql.port"),
		PgUsername:   viper.GetString("postgresql.username"),
		PgPassword:   viper.GetString("postgresql.password"),
		PgDatabase:   viper.GetString("postgresql.database"),
		PgSslMode:    viper.GetString("postgresql.sslmode"),
		PgTimezone:   viper.GetString("postgresql.timezone"),
		JwtKey:       viper.GetString("jwt.key"),
		JwtExpired:   viper.GetInt("jwt.expired"),
		JwtRefresh:   viper.GetInt("jwt.refresh"),

		// ADDITIONAL
		SMTPServer:   viper.GetString("email.SMTPServer"),
		SMTPPort:     viper.GetInt("email.SMTPPort"),
		SMTPUsername: viper.GetString("email.SMTPUsername"),
		SMTPPassword: viper.GetString("email.SMTPPassword"),
	}

	return cfg

}

var config = load()

func Get() *Config {
	return &config
}

func CheckConfig() merror.Error {
	// Check for required fields
	requiredFields := []string{
		"app.mode", "app.debug", "app.name", "app.url", "app.port", "app.api_key",
		"postgresql.hostname", "postgresql.port",
		"postgresql.username", "postgresql.password", "postgresql.database",
		"postgresql.sslmode", "postgresql.timezone",
	}

	for _, field := range requiredFields {
		if !viper.IsSet(field) || viper.GetString(field) == "" {
			return merror.RecordError(errors.New(fmt.Sprintf("%s is missing or empty", field)), http.StatusInternalServerError, "Unable to read Config")
		}
	}
	return merror.Error{}
}
