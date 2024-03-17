package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func bashExec(pathToScript string, parameters ...string) ([]byte, error) {

	args := append([]string{pathToScript}, parameters...)

	output, err := exec.Command("bash", args...).Output()

	return output, err
}

func getNume(c *gin.Context) {

	output, err := bashExec("/var/lib/licenta/api-licenta/get_nume_produs_sistem.sh")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getSerial(c *gin.Context, rootpass string) {

	output, err := bashExec("/var/lib/licenta/api-licenta/get_numar_serial_sistem.sh", rootpass)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getFurnizor(c *gin.Context) {

	output, err := bashExec("/var/lib/licenta/api-licenta/get_furnizor_sistem.sh")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getProcesor(c *gin.Context) {

	output, err := bashExec("/var/lib/licenta/api-licenta/get_procesor_sistem.sh")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getPlaciRetea(c *gin.Context) {

	output, err := bashExec("/var/lib/licenta/api-licenta/get_placi_retea.sh")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	data := strings.Split(string(output), "\n")

	c.JSON(http.StatusOK, data[:len(data)-1])
}

func getStarePlacaRetea(c *gin.Context) {

	param := c.Query("param")

	output, err := bashExec("/var/lib/licenta/api-licenta/get_stare_placa_retea.sh", param)
	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getDateTransmisePlacaRetea(c *gin.Context) {

	param := c.Query("param")

	output, err := bashExec("/var/lib/licenta/api-licenta/get_date_transmise_placa_retea.sh", param)
	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getDateReceptionatePlacaRetea(c *gin.Context) {

	param := c.Query("param")

	output, err := bashExec("/var/lib/licenta/api-licenta/get_date_receptionate_placa_retea.sh", param)
	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getDateAruncatePlacaRetea(c *gin.Context) {

	param := c.Query("param")

	output, err := bashExec("/var/lib/licenta/api-licenta/get_date_aruncate_placa_retea.sh", param)
	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func loadDB(c *gin.Context, db *sql.DB, rootpass string) {

	err := LoadDatabase(db, rootpass)

	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to load the database")
		return
	}

	c.String(http.StatusOK, "%s", "Database Loaded")
}

func triggerLoadDB(c *gin.Context, rootpass string, IPenv string) {

	_, err := bashExec("/var/lib/licenta/api-licenta/update_db.sh", rootpass, IPenv)

	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to trigger the load database crontab job")
		return
	}

	c.String(http.StatusOK, "%s", "Database Crontab Activated")
}
