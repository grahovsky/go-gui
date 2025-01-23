package models

type KeyConfig struct {
	Value string `yaml:"value"`
	Delay string `yaml:"delay"`
	Clip  string `yaml:"clip"`
}

type Config struct {
	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`
	Window struct {
		Name      string `yaml:"name"`
		ActivePid int
	} `yaml:"window"`
	HotKey string      `yaml:"hotkey"`
	Keys   []KeyConfig `yaml:"keys"`
}
