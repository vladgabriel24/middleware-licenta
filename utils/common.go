package utils

import "os/exec"

func BashExec(pathToScript string, parameters ...string) ([]byte, error) {

	args := append([]string{pathToScript}, parameters...)

	output, err := exec.Command("bash", args...).Output()

	return output, err
}
