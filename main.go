package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"github.com/tuannkhoi/PB-GraphQL-Proxy/graph"
	"github.com/tuannkhoi/PB-GraphQL-Proxy/graph/generated"
)

// graphqlHandler is the GraphQL endpoint.
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// playgroundHandler is the GraphQL playground.
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	router := gin.Default()

	router.GET("/", playgroundHandler())
	router.POST("/query", graphqlHandler())

	log.Info("Starting Gin on port 5000")
	if err := router.Run(":5000"); err != nil {
		log.Fatalf("got error running Gin: %v", err)
	}
}
