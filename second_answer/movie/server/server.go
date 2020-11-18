package main

import (
	"context"
	"flag"
	"fmt"
	proto "github.com/muhammadaser/stockbit_proto/movie"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 8081, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterRouteMovieServer(grpcServer, &movieServer{})
	log.Println("listen server to port " + fmt.Sprintf("localhost:%d", *port))
	grpcServer.Serve(lis)

}

type movieServer struct {
	proto.UnimplementedRouteMovieServer
}

func (s *movieServer) GetMovies(ctx context.Context, params *proto.MovieParams) (*proto.MoviesRes, error) {
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

func (s *movieServer) GetMovie(context.Context, *proto.SingleMovieParams) (*proto.Movie, error) {

	movie := proto.Movie{
		ImdbId:     "tt0371746",
		ImdbRating: "7.9",
		ImdbVotes:  "927,094",
		Title:      "Iron Man",
	}

	return &movie, nil
}
