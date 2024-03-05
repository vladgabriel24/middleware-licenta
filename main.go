package main

import (
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

	output, err := bashExec("/var/lib/licenta/api/get_nume_produs_sistem.sh")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getSerial(c *gin.Context) {

	output, err := bashExec("/var/lib/licenta/api/get_numar_serial_sistem.sh")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getFurnizor(c *gin.Context) {

	output, err := bashExec("/var/lib/licenta/api/get_furnizor_sistem.sh")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getProcesor(c *gin.Context) {

	output, err := bashExec("/var/lib/licenta/api/get_procesor_sistem.sh")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getPlaciRetea(c *gin.Context) {

	output, err := bashExec("/var/lib/licenta/api/get_placi_retea.sh")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	data := strings.Split(string(output), "\n")

	c.JSON(http.StatusOK, data[:len(data)-1])
}

func getStarePlacaRetea(c *gin.Context) {

	param := c.Query("param")

	output, err := bashExec("/var/lib/licenta/api/get_stare_placa_retea.sh", param)
	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getDateTransmisePlacaRetea(c *gin.Context) {

	param := c.Query("param")

	output, err := bashExec("/var/lib/licenta/api/get_date_transmise_placa_retea.sh", param)
	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getDateReceptionatePlacaRetea(c *gin.Context) {

	param := c.Query("param")

	output, err := bashExec("/var/lib/licenta/api/get_date_receptionate_placa_retea.sh", param)
	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getDateAruncatePlacaRetea(c *gin.Context) {

	param := c.Query("param")

	output, err := bashExec("/var/lib/licenta/api/get_date_aruncate_placa_retea.sh", param)
	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func main() {
	router := gin.Default()
	router.GET("/get-nume", getNume)
	router.GET("/get-serial", getSerial)
	router.GET("/get-furnizor", getFurnizor)
	router.GET("/get-procesor", getProcesor)
	router.GET("/get-placi-retea", getPlaciRetea)
	router.GET("/get-stare-placa-retea", getStarePlacaRetea)
	router.GET("/get-date-transmise-placa-retea", getDateTransmisePlacaRetea)
	router.GET("/get-date-receptionate-placa-retea", getDateReceptionatePlacaRetea)
	router.GET("/get-date-aruncate-placa-retea", getDateAruncatePlacaRetea)

	router.Run("localhost:8080")
}
