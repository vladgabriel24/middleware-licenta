package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
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

func getUtilizareDisk(c *gin.Context, rootpass string) {

	output, err := bashExec("/var/lib/licenta/api-licenta/get_utilizare_disk.sh", rootpass)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	output_lines := strings.Split(string(output), "\n")

	result := make(map[string][4]string)

	for i := 1; i < len(output_lines)-1; i++ {
		key := strings.Split(output_lines[i], " ")[2]

		tmpval_used := strings.Split(output_lines[i], " ")[0]
		tmpval_avail := strings.Split(output_lines[i], " ")[1]

		value_used, err_used := strconv.ParseFloat(tmpval_used[:len(tmpval_used)-1], 64)
		if err_used != nil {
			fmt.Println("Error at unit conversion for used variable disk usage")
		}

		value_avail, err_avail := strconv.ParseFloat(tmpval_avail[:len(tmpval_avail)-1], 64)
		if err_avail != nil {
			fmt.Println("Error at unit conversion for available variable disk usage")
		}

		unit_used := tmpval_used[len(tmpval_used)-1:]
		unit_avail := tmpval_avail[len(tmpval_avail)-1:]

		if unit_used == "K" {
			value_used = value_used / (1024 * 1024)
		}

		if unit_used == "M" {
			value_used = value_used / 1024
		}

		if unit_used == "T" {
			value_used = value_used * 1024
		}

		if unit_avail == "K" {
			value_avail = value_avail / (1024 * 1024)
		}

		if unit_avail == "M" {
			value_avail = value_avail / 1024
		}

		if unit_avail == "T" {
			value_avail = value_avail * 1024
		}

		value_free := value_avail - value_used

		percentage_used := (value_used / value_avail) * 100

		value := [4]string{fmt.Sprintf("%.2g", value_avail) + "G",
			fmt.Sprintf("%.2g", value_used) + "G",
			fmt.Sprintf("%.2g", value_free) + "G",
			fmt.Sprintf("%.2g", percentage_used) + "%"}

		result[string(key)] = value
	}

	c.JSON(http.StatusOK, result)
}

func getUtilizareRAM(c *gin.Context) {

	output, err := bashExec("/var/lib/licenta/api-licenta/get_utilizare_ram.sh")
	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	output_lines := strings.Split(string(output), "\n")

	tmpval_total, err_total := strconv.ParseFloat(strings.Split(output_lines[0], " ")[1], 64)
	if err_total != nil {
		fmt.Println("Error at conversion the total value for RAM utilization")
	}

	tmpval_avail, err_avail := strconv.ParseFloat(strings.Split(output_lines[1], " ")[1], 64)
	if err_avail != nil {
		fmt.Println("Error at conversion the available value for RAM utilization")
	}

	tmpval_used := tmpval_total - tmpval_avail

	result := map[string]string{
		"Free":  fmt.Sprintf("%.2g", tmpval_avail/(1024*1024)) + "G",
		"Total": fmt.Sprintf("%.2g", tmpval_total/(1024*1024)) + "G",
		"Used":  fmt.Sprintf("%.2g", tmpval_used/(1024*1024)) + "G",
	}

	c.JSON(http.StatusOK, result)
}

func getUtilizareCPU(c *gin.Context) {

	output, err := bashExec("/var/lib/licenta/api-licenta/get_utilizare_cpu.sh")
	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to execute script")
		return
	}

	output_lines := strings.Split(string(output), "\n")
	tmp_avgvals := strings.Split(output_lines[0], ":")[1]

	result := map[string]string{
		"15min":        strings.Split(tmp_avgvals, " ")[2],
		"1min":         strings.Split(tmp_avgvals, " ")[0],
		"5min":         strings.Split(tmp_avgvals, " ")[1],
		"noProcessors": strings.Split(output_lines[1], ":")[1],
	}

	c.JSON(http.StatusOK, result)
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
