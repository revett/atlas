package config

import "errors"

var ErrContextConfigValueIsNotConfigType = errors.New(
	"context config value cannot be cast to config type",
)
