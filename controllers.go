package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"middleware/web-service-gin/utils"

	"github.com/gin-gonic/gin"
)

func getNume(c *gin.Context) {

	output, err := utils.Nume()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getSerial(c *gin.Context, rootpass string) {

	output, err := utils.Serial(rootpass)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getFurnizor(c *gin.Context) {

	output, err := utils.Furnizor()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getProcesor(c *gin.Context) {

	output, err := utils.Procesor()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getPlaciRetea(c *gin.Context) {

	output, err := utils.PlaciRetea()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.JSON(http.StatusOK, output)
}

func getStarePlacaRetea(c *gin.Context) {

	param := c.Query("param")

	output, err := utils.StarePlacaRetea(param)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getDateTransmisePlacaRetea(c *gin.Context) {

	param := c.Query("param")

	output, err := utils.DateTransmisePlacaRetea(param)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getDateReceptionatePlacaRetea(c *gin.Context) {

	param := c.Query("param")

	output, err := utils.DateReceptionatePlacaRetea(param)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getDateAruncatePlacaRetea(c *gin.Context) {

	param := c.Query("param")

	output, err := utils.DateAruncatePlacaRetea(param)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getUtilizareDisk(c *gin.Context, rootpass string) {

	output, err := utils.UtilizareDisk(rootpass)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.JSON(http.StatusOK, output)
}

func getUtilizareRAM(c *gin.Context) {

	output, err := utils.UtilizareRAM()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.JSON(http.StatusOK, output)
}

func getUtilizareCPU(c *gin.Context) {

	output, err := utils.UtilizareCPU()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.JSON(http.StatusOK, output)
}

func loadDB(c *gin.Context, db *sql.DB, rootpass string) {

	err := utils.LoadDatabase(db, rootpass)

	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to load the database")
		return
	}

	c.String(http.StatusOK, "%s", "Database Loaded")
}

func triggerLoadDB(c *gin.Context, rootpass string, IPenv string) {

	_, err := utils.BashExec("/var/lib/licenta/api-licenta/update_db.sh", rootpass, IPenv)

	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to trigger the load database crontab job")
		return
	}

	c.String(http.StatusOK, "%s", "Database Crontab Activated")
}
