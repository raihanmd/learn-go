package test

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var config = viper.New()

func TestJson(t *testing.T) {
	config.SetConfigFile("./config/config.json")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	t.Log(config.AllKeys())
	t.Log(config.AllSettings())

	assert.Equal(t, "go-viper", config.GetString("app.name"))
	assert.Equal(t, "viper", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 5432, config.GetInt("database.port"))
}

func TestYaml(t *testing.T) {
	config.SetConfigFile("./config/config.yaml")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	t.Log(config.AllKeys())
	t.Log(config.AllSettings())

	assert.Equal(t, "go-viper", config.GetString("app.name"))
	assert.Equal(t, "viper", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 5432, config.GetInt("database.port"))
}

func TestEnvFile(t *testing.T) {
	config.SetConfigFile("./config/config.env")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	t.Log(config.AllKeys())
	t.Log(config.AllSettings())

	assert.Equal(t, "go-viper", config.GetString("app_name"))
	assert.Equal(t, "viper", config.GetString("app_author"))
	assert.Equal(t, "localhost", config.GetString("database_host"))
	assert.Equal(t, 5432, config.GetInt("database_port"))
}

func TestEnv(t *testing.T) {
	t.Log("GO_ENV:", os.Getenv("GO_ENV"))
	if os.Getenv("GO_ENV") != "production" {
		config.SetConfigFile("./config/config.env")
	} else {
		config.AutomaticEnv()
	}

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "go-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "viper", config.GetString("APP_AUTHOR"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 5432, config.GetInt("DATABASE_PORT"))

	assert.Equal(t, "Hello", config.GetString("FROM_ENV"))
}
