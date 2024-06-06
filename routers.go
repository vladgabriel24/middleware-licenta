package main

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initRouters(ip string, port int, rootpass string, database *sql.DB) {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/api/get-nume", getNume)
	router.GET("/api/get-serial", func(ctx *gin.Context) {
		getSerial(ctx, rootpass)
	})
	router.GET("/api/get-furnizor", getFurnizor)
	router.GET("/api/get-procesor", getProcesor)
	router.GET("/api/get-placi-retea", getPlaciRetea)
	router.GET("/api/get-stare-placa-retea", getStarePlacaRetea)
	router.GET("/api/get-date-transmise-placa-retea", getDateTransmisePlacaRetea)
	router.GET("/api/get-date-receptionate-placa-retea", getDateReceptionatePlacaRetea)
	router.GET("/api/get-date-aruncate-placa-retea", getDateAruncatePlacaRetea)
	router.GET("/api/get-date-retea", getDateRetea)
	router.GET("/api/get-utilizare-disk", func(ctx *gin.Context) {
		getUtilizareDisk(ctx, rootpass)
	})
	router.GET("/api/get-utilizare-RAM", getUtilizareRAM)
	router.GET("/api/get-utilizare-CPU", getUtilizareCPU)
	router.POST("/api/load-db", func(ctx *gin.Context) {
		loadDB(ctx, database, rootpass)
	})

	router.Run(ip + ":" + strconv.Itoa(port))
}
