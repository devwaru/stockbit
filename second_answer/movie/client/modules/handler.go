package modules

import (
	"context"
	proto "github.com/muhammadaser/stockbit_proto/movie"
	"log"
	"net/http"
)

func (m *Module) HandleMovies(w http.ResponseWriter, h *http.Request) {
	movies, err := m.MovieClient.GetMovies(context.Background(), &proto.MovieParams{SearchWord: "batman", Page: 1})
	if err != nil {
		log.Println(err)
	}

	toJSON(w, http.StatusOK, movies)
}

func (m *Module) HandleMovie(w http.ResponseWriter, h *http.Request) {

	toJSON(w, http.StatusOK, "ok")

}
