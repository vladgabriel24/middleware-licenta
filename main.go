package main

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func getNume(c *gin.Context) {

	cmd := exec.Command("bash", "/var/lib/licenta/api/get_nume_produs_sistem.sh")
	output, err := cmd.Output()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getSerial(c *gin.Context) {

	cmd := exec.Command("bash", "/var/lib/licenta/api/get_numar_serial_sistem.sh")
	output, err := cmd.Output()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getFurnizor(c *gin.Context) {

	cmd := exec.Command("bash", "/var/lib/licenta/api/get_furnizor_sistem.sh")
	output, err := cmd.Output()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	c.String(http.StatusOK, "%s", output)
}

func getProcesor(c *gin.Context) {

	cmd := exec.Command("bash", "/var/lib/licenta/api/get_procesor_sistem.sh")
	output, err := cmd.Output()
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

	router.Run("localhost:8080")
}
