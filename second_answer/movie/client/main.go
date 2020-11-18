package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/muhammadaser/stockbit/second_answer/config"
	"github.com/muhammadaser/stockbit/second_answer/modules"
	"log"
	"net/http"
)

func main() {
	var env string
	flag.StringVar(&env, "env", "dev", "set env here etc : dev,stg,prd")
	flag.Parse()

	cfg, err := config.InitConfig(env)
	if err != nil {
		log.Fatal("error init config: ", err.Error())
	}

	log.Printf("given config: %+v", cfg)

	flag.Parse()

	m := modules.New(cfg)

	// http routing
	r := mux.NewRouter()
	r.HandleFunc("/v1/movies", m.HandleMovies).Methods(http.MethodPost)
	r.HandleFunc("/v1/movie/{movie_id}", m.HandleMovie).Methods(http.MethodPost)

}
