package modules

import (
	"github.com/muhammadaser/stockbit/second_answer/config"
)

type Module struct {
	Cfg *config.Config
}

func New(cfg *config.Config) *Module {
	module := Module{}
	module.Cfg = cfg

	return &module
}
