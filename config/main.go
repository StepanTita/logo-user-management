package config

import (
	"fmt"
	_ "github.com/lib/pq" // here
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config interface {
	Listener() string
	Logger
	Databaser
}

// main system configuration
type config struct {
	Addr     string `yaml:"addr"`
	Log      string `yaml:"log"`
	Database struct {
		URL    string `yaml:"url"`
		Method string `yaml:"migrate"`
	} `yaml:"db"`

	Logger
	Databaser
}

// initializes New system configuration object
func New(path string) Config {
	cfg := config{}

	yamlConfig, err := ioutil.ReadFile(path)
	if err != nil {
		panic(errors.New(fmt.Sprintf("failed to read config: %s", path)))
	}

	err = yaml.Unmarshal(yamlConfig, &cfg)
	if err != nil {
		panic(errors.New(fmt.Sprintf("failed to unmarshal config: %s", path)))
	}

	if port := os.Getenv("PORT"); port != "" {
		cfg.Addr = ":" + port
	}
	cfg.Logger = NewLogger(cfg.Log)
	cfg.Databaser = NewDatabaser(cfg.Database.URL, cfg.Database.Method, cfg.Logger.Logging())

	return &cfg
}

func (c *config) Listener() string {
	return c.Addr
}
