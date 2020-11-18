package modules

import (
	"github.com/muhammadaser/stockbit/second_answer/movie/server/config"
	proto "github.com/muhammadaser/stockbit_proto/movie"
)

type Module struct {
	proto.UnimplementedRouteMovieServer
	Cfg *config.Config
}

func New(cfg *config.Config) *Module {
	module := Module{}
	module.Cfg = cfg

	return &module
}
