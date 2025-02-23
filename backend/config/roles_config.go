package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Permissions []string

type Role struct {
	Permissions Permissions `yaml:"permissions"`
}

type RolesConfig struct {
	Roles map[string]Role `yaml:"roles"`
}

var rolesConfiguration *RolesConfig

func LoadRolesConfig(filepath string) (*RolesConfig, error) {
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var config RolesConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	rolesConfiguration = &config

	return &config, nil
}

func GetRolesConfig() *RolesConfig {
	return rolesConfiguration
}
