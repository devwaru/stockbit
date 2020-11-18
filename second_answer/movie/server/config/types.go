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
