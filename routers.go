package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func initRouters(ip string, port int, rootpass string) {

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

	router.Run(ip + ":" + strconv.Itoa(port))
}
