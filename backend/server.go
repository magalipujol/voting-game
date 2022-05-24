package main

import (
	"api/graph"
	"api/graph/generated"
	"api/repository"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

func main() {
	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
		// AllowedMethods:   []string{"GET", "POST"},
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		PlayerRepository: repository.NewPlayerRepository(),
		RoomRepository:   repository.NewRoomRepository(),
	}}))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/", playground.Handler("voting game", "/graphql"))
	router.Handle("/graphql", srv)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
