package main

import (
	"graphql/graph"
	"graphql/graph/dbutil"
	"graphql/graph/generated"
	"graphql/mock"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
)

func main() {
	db := dbutil.NewDB()

	mock.InsertMockData(db.Conn)

	resolver := &graph.Resolver{DB: db}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Println("Server running at http://localhost:8081/")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
