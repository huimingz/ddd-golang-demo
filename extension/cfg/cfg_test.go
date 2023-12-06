package cfg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type config struct {
	t *testing.T

	after bool
}

func (c *config) BeforeLoad() {
	assert.False(c.t, c.after)
	c.after = true
}

func (c *config) AfterLoad() {
	assert.True(c.t, c.after)
}

type vali struct {
	Address string `validate:"lt=10"`                // 字符串长度<10
	Cloud   string `validate:"oneof=ali tx private"` // 3个字符串中的一个
}

func (cfg *config) Validate() error {
	return fmt.Errorf("check error")
}

func TestValidate(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	v := &vali{
		Address: "123456789",
		Cloud:   "ali",
	}
	err := validate(v)
	is.Nil(err)

	v1 := &vali{
		Address: "12345678900",
		Cloud:   "tx",
	}
	err = validate(v1)
	is.NotNil(err)

	c := &config{}
	err = customValidate(c)

	is.NotNil(err)
	is.Equal("check error", err.Error())
}

func TestHook(t *testing.T) {
	t.Parallel()
	c := &config{t: t}
	Load(c)
}
