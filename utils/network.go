package utils

import (
	"strings"
)

func PlaciRetea() ([]string, error) {

	output, err := BashExec("/var/lib/licenta/api-licenta/get_placi_retea.sh")
	if err != nil {
		return []string{}, err
	}

	data := strings.Split(string(output), "\n")

	return data[:len(data)-1], nil
}

func StarePlacaRetea(param string) (string, error) {

	output, err := BashExec("/var/lib/licenta/api-licenta/get_stare_placa_retea.sh", param)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func DateTransmisePlacaRetea(param string) (string, error) {

	output, err := BashExec("/var/lib/licenta/api-licenta/get_date_transmise_placa_retea.sh", param)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func DateReceptionatePlacaRetea(param string) (string, error) {

	output, err := BashExec("/var/lib/licenta/api-licenta/get_date_receptionate_placa_retea.sh", param)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func DateAruncatePlacaRetea(param string) (string, error) {

	output, err := BashExec("/var/lib/licenta/api-licenta/get_date_aruncate_placa_retea.sh", param)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
