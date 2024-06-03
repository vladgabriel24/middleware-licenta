package main

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
)

func initRouters(ip string, port int, rootpass string, database *sql.DB) {

	router := gin.Default()
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
	router.POST("/api/trigger-cron_db", func(ctx *gin.Context) {
		triggerLoadDB(ctx, rootpass, ip)
	})

	router.Run(ip + ":" + strconv.Itoa(port))
}
