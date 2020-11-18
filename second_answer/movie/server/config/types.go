package config

type Server struct {
	Port               int
	DefaultHttpTimeout int    `toml:"http_timeout"`
	MovieServerAddr    string `toml:"movie_server_addr"`
}

type DB struct {
	DSN         string `toml:"dsn"`
	MaxConn     int    `toml:"max_conn"`
	MaxIdleConn int    `toml:"max_idle_conn"`
}

// for now, just put api key on config file, next improvement, need to move to other place
type Omdb struct {
	Host   string `toml:"host"`
	ApiKey string `toml:"api_key"`
}
