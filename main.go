package main

import (
	"net/http"
	"os"

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

	router.Use(CORSMiddleware())

	router.GET("/", playgroundHandler())
	router.POST("/query", graphqlHandler())

	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	}

	log.Infof("Starting GraphQL Proxy at localhost%s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("got error running GraphQL Proxy: %v", err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")

		// handle CORS preflight requests
		// https://developer.mozilla.org/en-US/docs/Glossary/Preflight_request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next()
		}
	}
}

