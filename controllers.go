package main

import (
	"database/sql"
	"net/http"

	"middleware/utils"

	"github.com/gin-gonic/gin"
)

func getNume(c *gin.Context) {

	output, err := utils.Nume()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, map[string]error{"Failed to execute the get_nume_produs_sistem script": err})
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getSerial(c *gin.Context, rootpass string) {

	output, err := utils.Serial(rootpass)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, map[string]error{"Failed to execute the get_numar_serial_sistem script": err})
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getFurnizor(c *gin.Context) {

	output, err := utils.Furnizor()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, map[string]error{"Failed to execute the get_furnizor_sistem script": err})
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getProcesor(c *gin.Context) {

	output, err := utils.Procesor()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, map[string]error{"Failed to execute the get_procesor_sistem script": err})
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getPlaciRetea(c *gin.Context) {

	output, err := utils.PlaciRetea()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, map[string]error{"Failed to execute the get_placi_retea script": err})
		return
	}

	c.JSON(http.StatusOK, output)
}

func getStarePlacaRetea(c *gin.Context) {

	param := c.Query("param")

	output, err := utils.StarePlacaRetea(param)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, map[string]error{"Failed to execute the get_stare_placa_retea script": err})
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getDateTransmisePlacaRetea(c *gin.Context) {

	param := c.Query("param")

	output, err := utils.DateTransmisePlacaRetea(param)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, map[string]error{"Failed to execute the get_date_transmise_placa_retea script": err})
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getDateReceptionatePlacaRetea(c *gin.Context) {

	param := c.Query("param")

	output, err := utils.DateReceptionatePlacaRetea(param)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, map[string]error{"Failed to execute the get_date_receptionate_placa_retea script": err})
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getDateAruncatePlacaRetea(c *gin.Context) {

	param := c.Query("param")

	output, err := utils.DateAruncatePlacaRetea(param)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, map[string]error{"Failed to execute the get_date_aruncate_placa_retea script": err})
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getUtilizareDisk(c *gin.Context, rootpass string) {

	output, err := utils.UtilizareDisk(rootpass)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, map[string]error{"Failed to execute the get_utilizare_disk script": err})
		return
	}

	c.JSON(http.StatusOK, output)
}

func getUtilizareRAM(c *gin.Context) {

	output, err := utils.UtilizareRAM()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, map[string]error{"Failed to execute the get_utilizare_ram script": err})
		return
	}

	c.JSON(http.StatusOK, output)
}

func getUtilizareCPU(c *gin.Context) {

	output, err := utils.UtilizareCPU()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, map[string]error{"Failed to execute the get_utilizare_cpu script": err})
		return
	}

	c.JSON(http.StatusOK, output)
}

func loadDB(c *gin.Context, db *sql.DB, rootpass string) {

	err := utils.LoadDatabase(db, rootpass)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, map[string]error{"Failed to load the database": err})
		return
	}

	c.String(http.StatusOK, "%s", "Database Loaded")
}

func triggerLoadDB(c *gin.Context, rootpass string, IPenv string) {

	output, err := utils.TriggerLoadCrontab(rootpass, IPenv)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, map[string]error{"Failed to trigger the load database crontab job": err})
		return
	}

	c.String(http.StatusOK, "%s", output)
}
