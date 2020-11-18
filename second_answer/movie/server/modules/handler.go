package modules

import (
	"context"
	proto "github.com/muhammadaser/stockbit_proto/movie"
)

func (m *Module) GetMovies(ctx context.Context, params *proto.MovieParams) (*proto.MoviesRes, error) {
	movie := proto.Movie{
		ImdbId:     "tt0371746",
		ImdbRating: "7.9",
		ImdbVotes:  "927,094",
		Title:      "Iron Man",
	}
	var movies []*proto.Movie
	movies = append(movies, &movie)

	res := proto.MoviesRes{
		Movies: movies,
	}

	return &res, nil
}

func (m *Module) GetMovie(context.Context, *proto.SingleMovieParams) (*proto.Movie, error) {

	movie := proto.Movie{
		ImdbId:     "tt0371746",
		ImdbRating: "7.9",
		ImdbVotes:  "927,094",
		Title:      "Iron Man",
	}

	return &movie, nil
}
