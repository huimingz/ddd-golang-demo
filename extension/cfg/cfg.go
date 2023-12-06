package cfg

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

// IValidate is the interface that check config options
type IValidate interface {
	Validate() error
}

type IHookBeforeLoad interface {
	BeforeLoad()
}

type IHookAfterLoad interface {
	AfterLoad()
}

// Options is the config options
type Options struct {
	ConfigPath   string
	ConfigType   string
	ConfigName   string
	OpenDefault  bool // 是否开启默认值读取: "default" tag
	OpenValidate bool // 是否开启配置校验: "validate" tag
}

func defaultOption() (*Options, error) {
	workDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("get work dir error when set default config options: %s", err.Error())
	}
	dopt := &Options{
		ConfigPath:   workDir,
		ConfigType:   "yaml",
		ConfigName:   "config",
		OpenDefault:  true,
		OpenValidate: true,
	}

	return dopt, nil
}

// Load load config
// if don't pass opts params, will use default options:
//   - ConfigPath: current work dir
//   - ConfigType: "yaml"
//   - ConfigName: "config"
//   - OpenDefault: true
//   - OpenValidate: true
func Load(v interface{}, opts ...*Options) error {
	return load(v, opts...)
}

// MustLoad load config, if error occur, panic.
// if don't pass opts params, will use default options:
//   - ConfigPath: current work dir
//   - ConfigType: "yaml"
//   - ConfigName: "config"
//   - OpenDefault: true
//   - OpenValidate: true
func MustLoad(v interface{}, opts ...*Options) {
	if err := load(v, opts...); err != nil {
		panic(err)
	}
}

func load(v interface{}, opts ...*Options) error {
	var ropt *Options
	if len(opts) > 0 {
		ropt = opts[0]
	}

	if hooker, ok := v.(IHookBeforeLoad); ok {
		hooker.BeforeLoad()
	}

	if ropt == nil {
		var err error
		ropt, err = defaultOption()
		if err != nil {
			return err
		}
	}

	viper.SetConfigName(ropt.ConfigName)
	viper.SetConfigType(ropt.ConfigType)
	viper.AddConfigPath(ropt.ConfigPath)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read config error: %s", err.Error())
	}

	if ropt.OpenDefault {
		defaults.SetDefaults(v)
	}

	if err := viper.Unmarshal(v); err != nil {
		return fmt.Errorf("init configuration failed: %v", err)
	}

	if ropt.OpenValidate {
		err := validate(v)
		if err != nil {
			return err
		}
	}

	if err := customValidate(v); err != nil {
		return nil
	}

	if hooker, ok := v.(IHookAfterLoad); ok {
		hooker.AfterLoad()
	}

	return nil
}

func validate(v interface{}) error {
	vali := validator.New()

	err := vali.Struct(v)
	if err != nil {
		return err
	}

	return nil
}

func customValidate(v interface{}) error {
	if checker, ok := v.(IValidate); ok {
		if err := checker.Validate(); err != nil {
			return err
		}
	}

	return nil
}
