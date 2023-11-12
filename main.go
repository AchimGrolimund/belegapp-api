package main

import (
	"beleg-app/api/handler"
	"beleg-app/api/repository"
	"beleg-app/api/service"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbURL := "postgres://root:root@localhost:5432/beleg_db"
	dbPool, err := connectToDB(dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbPool.Close()

	// Initialisieren des Beleg Repositories
	belegRepo := repository.NewBelegRepository(dbPool)

	// Initialisieren des Beleg Services
	belegService := service.NewBelegService(belegRepo)

	// Initialisieren des Beleg Handlers
	belegHandler := handler.NewBelegHandler(belegService)

	// Einrichten des Gin-Routers
	r := gin.Default()
	r.SecureJsonPrefix(")]}',\n")

	// Einfache Ping-Route als Health Check
	r.GET("/ping", func(c *gin.Context) {
		c.SecureJSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Setup der API-Routen
	setupRoutes(r, belegHandler)

	// Starten des Servers
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v\n", err)
	}
}

func connectToDB(dbURL string) (*pgxpool.Pool, error) {
	connPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}
	if err = connPool.Ping(context.Background()); err != nil {
		return nil, err
	}
	return connPool, nil
}

func setupRoutes(router *gin.Engine, belegHandler *handler.BelegHandler) {
	// Gruppe für API Version 1
	v1 := router.Group("/v1/")
	{
		v1.GET("/beleg/:id", belegHandler.GetBelegById)
		v1.GET("/beleg", belegHandler.GetAllBelege)
		v1.POST("/beleg", belegHandler.CreateBeleg)
		v1.DELETE("/beleg/:id", belegHandler.DeleteBelegById)
	}
}
