package utils

func Nume() (string, error) {

	output, err := BashExec("/var/lib/licenta/api-licenta/get_nume_produs_sistem.sh")
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func Serial(rootpass string) (string, error) {

	output, err := BashExec("/var/lib/licenta/api-licenta/get_numar_serial_sistem.sh", rootpass)
	if err != nil {
		return "", err
	}

	return string(output), err
}

func Furnizor() (string, error) {

	output, err := BashExec("/var/lib/licenta/api-licenta/get_furnizor_sistem.sh")
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func Procesor() (string, error) {

	output, err := BashExec("/var/lib/licenta/api-licenta/get_procesor_sistem.sh")
	if err != nil {
		return "", err
	}

	return string(output[1:]), err
}
