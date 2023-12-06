# cfg

## Features
- [x] 使用 validator.v10 进行配置校验
- [x] 使用`default`tag指定默认值
- [x] 支持自定义校验接口
- [x] 支持 BeforeLoad 和 AfterLoad 钩子函数

## 默认参数
- 默认配置路径: 运行目录
- 默认文件名: config.yaml
- 默认开启默认配置、validator 校验

## Examples
### 配置struct示例
```golang
type Config struct {
    // 配置名：value
    // 校验规则：值只能为json或 text
    // 默认值：text
    Value string `mapstructure:"value" validate:"oneof:json,text" default:"text"`
}
```

### 自定义校验规则
```golang
type Config struct {
    Value string `mapstructure:"value" validate:"oneof:json,text" default:"text"`
}

// 实现 cfg.IValidate 接口
func (c *Config) Validate() error {
    // do some validation here
}
```

### 钩子
```golang
type Config struct {
    Value string `mapstructure:"value" validate:"oneof:json,text" default:"text"`
}

// 实现 cfg.IHookBeforeLoad 接口
func (c *Config) BeforeLoad() error {
    // do something before load config file
}

// 实现 cfg.IHookAfterLoad 接口
func (c *Config) AfterLoad() error {
    // do something after load config file
}
```
