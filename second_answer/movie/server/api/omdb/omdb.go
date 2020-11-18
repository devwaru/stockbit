package omdb

import (
	"encoding/json"
	"fmt"
	"github.com/muhammadaser/stockbit/second_answer/movie/server/network"
	"log"
	"net/http"
)

func (o *Omdb) GetMovies(httpClient *http.Client, params MoviesRequest) (*MoviesResponse, error) {

	// query params
	queryParams := map[string]string{
		"apikey": o.Credential,
		"s":      params.SearchWord,
		"page":   fmt.Sprintf("%d", params.Page),
	}

	resByte, err := network.Do(httpClient, http.MethodGet, o.Host, queryParams, nil, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var res MoviesResponse

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &res, nil

}

func (o *Omdb) GetMovie(httpClient *http.Client, imdbId string) (*MovieResponse, error) {

	// query params
	queryParams := map[string]string{
		"apikey": o.Credential,
		"i":      imdbId,
	}

	resByte, err := network.Do(httpClient, http.MethodGet, o.Host, queryParams, nil, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var res MovieResponse

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &res, nil

}
