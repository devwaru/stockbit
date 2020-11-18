package main

import (
	"flag"
	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/muhammadaser/stockbit/second_answer/movie/client/config"
	"github.com/muhammadaser/stockbit/second_answer/movie/client/modules"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// logging : for now just use default system logging, will set letter

	// get env form parameter
	var env string
	flag.StringVar(&env, "env", "dev", "set env here eg : dev,stg,prd")
	flag.Parse()

	// init config
	cfg, err := config.InitConfig(env)
	if err != nil {
		log.Fatal("error init config: ", err.Error())
	}

	log.Printf("given config: %+v", cfg)

	flag.Parse()

	m := modules.New(cfg)

	// http routing
	r := mux.NewRouter()
	r.HandleFunc("/v1/movies", m.HandleMovies).Methods(http.MethodGet)
	r.HandleFunc("/v1/movie/{movie_id}", m.HandleMovie).Methods(http.MethodGet)

	// listen server
	log.Println("starting serve on ", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(cfg.Server.Port), gziphandler.GzipHandler(r)))

}
