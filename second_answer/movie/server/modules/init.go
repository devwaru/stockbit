package modules

import (
	"github.com/muhammadaser/stockbit/second_answer/movie/server/api/omdb"
	"github.com/muhammadaser/stockbit/second_answer/movie/server/config"
	"github.com/muhammadaser/stockbit/second_answer/movie/server/network"
	proto "github.com/muhammadaser/stockbit_proto/movie"
	"log"
	"net/http"
)

type Module struct {
	proto.UnimplementedRouteMovieServer
	Cfg        *config.Config
	HttpClient *http.Client
	OmdbApi    *omdb.Omdb
}

func New(cfg *config.Config) *Module {
	module := Module{}
	module.Cfg = cfg

	module.InitHttpClient()
	module.InitOmdbApi()

	return &module
}

func (m *Module) InitHttpClient() {
	hClient := network.New(m.Cfg.Server.DefaultHttpTimeout)
	log.Print("http client initiated")
	m.HttpClient = hClient
}

func (m *Module) InitOmdbApi() {
	omdbApi := omdb.InitOmdb(m.Cfg.Omdb.ApiKey, m.Cfg.Omdb.Host)
	log.Print("omdb api initiated")
	m.OmdbApi = omdbApi
}
