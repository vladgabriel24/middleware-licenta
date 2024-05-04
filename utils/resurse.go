package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func UtilizareDisk(rootpass string) (map[string][4]string, error) {

	output, err := BashExec("/var/lib/licenta/api-licenta/get_utilizare_disk.sh", rootpass)
	if err != nil {
		return nil, err
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

	return result, nil
}

func UtilizareRAM() (map[string]string, error) {

	output, err := BashExec("/var/lib/licenta/api-licenta/get_utilizare_ram.sh")
	if err != nil {
		return nil, err
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

	return result, err
}

func UtilizareCPU() (map[string]string, error) {

	output, err := BashExec("/var/lib/licenta/api-licenta/get_utilizare_cpu.sh")
	if err != nil {
		return nil, err
	}

	output_lines := strings.Split(string(output), "\n")
	tmp_avgvals := strings.Split(output_lines[0], ":")[1]

	result := map[string]string{
		"15min":        strings.Split(tmp_avgvals, " ")[2],
		"1min":         strings.Split(tmp_avgvals, " ")[0],
		"5min":         strings.Split(tmp_avgvals, " ")[1],
		"noProcessors": strings.Split(output_lines[1], ":")[1],
	}

	return result, nil
}
