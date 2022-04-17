package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var GET *Config = nil

type Config struct {
	LatestWorkspaces []string `yaml:"latest_workspaces"`
}

func (c *Config) Save() error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	if _, err = os.Stat(GetConfigDirectory()); os.IsNotExist(err) {
		err = os.MkdirAll(GetConfigDirectory(), 0755)
		if err != nil {
			return err
		}
	}

	return ioutil.WriteFile(filepath.Join(GetConfigDirectory(), "config.yaml"), data, 0644)
}

func Load() error {
	data, err := ioutil.ReadFile(filepath.Join(GetConfigDirectory(), "config.yaml"))
	if err != nil {
		return err
	}

	cfg := Config{}
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return err
	}

	GET = &cfg

	if err = GET.Save(); err != nil {
		return err
	}

	return nil
}

func GetConfigDirectory() string {
	cfgPath, err := os.UserConfigDir()
	if err != nil {
		cfgPath = "."
	}

	return filepath.Join(cfgPath, "gradient")
}
