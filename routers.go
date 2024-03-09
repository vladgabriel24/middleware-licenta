package main

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
)

func initRouters(ip string, port int, rootpass string, database *sql.DB) {

	router := gin.Default()
	router.GET("/get-nume", getNume)
	router.GET("/get-serial", func(ctx *gin.Context) {
		getSerial(ctx, rootpass)
	})
	router.GET("/get-furnizor", getFurnizor)
	router.GET("/get-procesor", getProcesor)
	router.GET("/get-placi-retea", getPlaciRetea)
	router.GET("/get-stare-placa-retea", getStarePlacaRetea)
	router.GET("/get-date-transmise-placa-retea", getDateTransmisePlacaRetea)
	router.GET("/get-date-receptionate-placa-retea", getDateReceptionatePlacaRetea)
	router.GET("/get-date-aruncate-placa-retea", getDateAruncatePlacaRetea)
	router.POST("/load-db", func(ctx *gin.Context) {
		loadDB(ctx, database, rootpass)
	})
	router.POST("/trigger-cron_db", func(ctx *gin.Context) {
		triggerLoadDB(ctx, rootpass)
	})

	router.Run(ip + ":" + strconv.Itoa(port))
}
