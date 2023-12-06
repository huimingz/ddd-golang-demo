package main

import (
	"fmt"

	"scan/extension/cfg"
)

var _ cfg.IHookAfterLoad = &Config{}
var _ cfg.IHookBeforeLoad = &Config{}
var _ cfg.IValidate = &Config{}

type Config struct {
	Switch bool `mapstructure:"switch"`
}

func (config *Config) Validate() error {
	fmt.Println("custom validate")
	return nil
}

func (config *Config) BeforeLoad() {
	fmt.Println("before load")
}

func (config *Config) AfterLoad() {
	fmt.Println("after load")
}

func main() {
	cfg.Load(&Config{})
}
