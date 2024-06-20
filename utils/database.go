package utils

import (
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func ConnectDB(user string, pass string, ip string, port int) (*sql.DB, error) {

	fmt.Println("Connecting to the database...")

	cfg_db := mysql.Config{
		User:                 user,
		Passwd:               pass,
		Net:                  "tcp",
		Addr:                 ip,
		DBName:               "metriciDB",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg_db.FormatDSN())
	if err != nil {
		fmt.Println("An Error appeared while opening the database...")
		return nil, err
	}

	errcon := db.Ping()
	if errcon != nil {
		fmt.Println("An Error appeared while connecting the database...")
		return nil, err
	}

	fmt.Println("Connected to the database.")
	return db, nil
}

func InitDB(user string, pass string, ip string, port int) *sql.DB {

	db, errdb := ConnectDB(user, pass, ip, port)
	if errdb != nil {
		log.Fatal(errdb)
	}
	// defer db.Close()

	return db

}

func LoadTblModel(db *sql.DB) (string, error) {

	// Preluam informatiile de sistem
	produs, errProdus := Nume()
	if errProdus != nil {
		return "", errProdus
	}
	fmt.Println(string(produs))

	_, errTblModel := db.Exec(
		`INSERT INTO
		tblModel (numeModel)
		VALUES (?)`,
		string(produs))

	if errTblModel != nil {
		if errTblModel.(*mysql.MySQLError).Number == 1062 {
			fmt.Println("Modelul se afla deja in baza de date")
		} else {
			return string(produs), errTblModel
		}
	}

	return string(produs), nil
}

func LoadTblProducator(db *sql.DB) (string, error) {

	// Preluam informatiile de sistem
	furnizor, errFurnizor := Furnizor()
	if errFurnizor != nil {
		return "", errFurnizor
	}
	fmt.Println(string(furnizor))

	_, errTblProducator := db.Exec(
		`INSERT INTO
		tblProducator (numeProducator)
		VALUES
		(?)`,
		string(furnizor))

	if errTblProducator != nil {
		if errTblProducator.(*mysql.MySQLError).Number == 1062 { // Identificator pentru incalcarea constrangerii de unique
			fmt.Println("Producatorul se afla deja in baza de date")
		} else {
			return string(furnizor), errTblProducator
		}
	}

	return string(furnizor), nil
}

func LoadTblProcesor(db *sql.DB) (string, error) {

	// Preluam informatiile de sistem
	procesor, errProcesor := Procesor()
	if errProcesor != nil {
		return "", errProcesor
	}
	fmt.Println(string(procesor))

	_, errTblProcesor := db.Exec(
		`INSERT INTO
		tblProcesor (numeProcesor)
		VALUES
		(?)`,
		string(procesor))

	if errTblProcesor != nil {
		if errTblProcesor.(*mysql.MySQLError).Number == 1062 {
			fmt.Println("Procesorul se afla deja in baza de date")
		} else {
			return string(procesor), errTblProcesor
		}
	}

	return string(procesor), nil
}

func LoadTblSerial(db *sql.DB, rootpass string) (string, error) {

	// Preluam informatiile de sistem
	serial, errSerial := Serial(rootpass)
	if errSerial != nil {
		return "", errSerial
	}
	fmt.Println(string(serial))

	_, errTblSerial := db.Exec(
		`INSERT INTO
		tblSerial (numarSerial)
		VALUES
		(?)`,
		string(serial))

	if errTblSerial != nil {
		if errTblSerial.(*mysql.MySQLError).Number == 1062 {
			fmt.Println("Numarul serial se afla deja in baza de date")
		} else {
			return string(serial), errTblSerial
		}
	}

	return string(serial), nil
}

func LoadTblSistem(db *sql.DB, rootpass string, procesor string, furnizor string, produs string, serial string) (string, error) {

	_, errTblSistem := db.Exec(
		`INSERT INTO
		tblSistem (numarSerial,modelProcesor,producatorSistem,modelSistem)
		VALUES (
			(
				SELECT idSerial
				FROM tblSerial
				WHERE numarSerial = ?
			),
			(
				SELECT idProcesor
				FROM tblProcesor
				WHERE numeProcesor = ?
			),
			(
				SELECT idProducator
				FROM tblProducator
				WHERE numeProducator = ?
			),
			(
				SELECT idModel
				FROM tblModel
				WHERE numeModel = ?
			)
		)`,
		string(serial), string(procesor), string(furnizor), string(produs))

	if errTblSistem != nil {
		if errTblSistem.(*mysql.MySQLError).Number == 1062 {
			fmt.Println("Sistemul se afla deja in baza de date")
		} else {
			return string(serial), errTblSistem
		}
	}

	return string(serial), nil
}

func LoadTblPlaciRetea(db *sql.DB, produs string, serial string) error {

	placi_retea, errNIC := PlaciRetea()
	if errNIC != nil {
		return errNIC
	}

	fmt.Println(len(placi_retea))

	for i := 0; i < len(placi_retea); i++ {

		fmt.Println(placi_retea[i])

		stare_nic, errStareNIC := StarePlacaRetea(placi_retea[i])
		if errStareNIC != nil {
			return errStareNIC
		}
		fmt.Println(string(stare_nic))

		var stare_nic_codata int

		if strings.Contains(string(stare_nic), "up") {
			stare_nic_codata = 1
		} else if strings.Contains(string(stare_nic), "down") {
			stare_nic_codata = 0
		} else {
			stare_nic_codata = 2
		}

		tx_nic, errTxNIC := DateTransmisePlacaRetea(placi_retea[i])
		if errTxNIC != nil {
			return errTxNIC
		}
		fmt.Println(string(tx_nic))

		rx_nic, errRxNIC := DateReceptionatePlacaRetea(placi_retea[i])
		if errRxNIC != nil {
			return errRxNIC
		}
		fmt.Println(string(rx_nic))

		dropped_nic, errDroppedNIC := DateAruncatePlacaRetea(placi_retea[i])
		if errDroppedNIC != nil {
			return errDroppedNIC
		}
		fmt.Println(string(dropped_nic))

		_, errTblPlaciRetea := db.Exec(
			`INSERT INTO
			tblPlaciRetea (
				modelSistem, 
				numarSerialSistem, 
				numePlacaRetea, 
				starePlacaRetea,
				pacheteAruncate,
				dateReceptionate,
				dateTransmise
			)
			VALUES (
				(
					SELECT idModel
					FROM tblModel
					WHERE numeModel = ?
				),
				(
					SELECT idSerial
					FROM tblSerial
					WHERE numarSerial = ?
				),
				?,
				?,
				?,
				?,
				?
			)`,
			string(produs), string(serial), string(placi_retea[i]), stare_nic_codata, string(dropped_nic), string(rx_nic), string(tx_nic))

		if errTblPlaciRetea != nil {
			return errTblPlaciRetea
		}
	}

	return nil
}

func LoadTblResurse(db *sql.DB, rootpass string, produs string, serial string) error {

	// used/available for Disk

	outputDisk, errDisk := UtilizareDisk(rootpass)
	if errDisk != nil {
		fmt.Println("Failed to execute the disk script for database load")
		return errDisk
	}

	tmpUsedDisk, errTmpUsed := strconv.ParseFloat(outputDisk["/"][1][:len(outputDisk["/"][1])-1], 64)
	if errTmpUsed != nil {
		fmt.Println("Error at conversion the \"Used\" value for Disk utilization at database load")
	}

	tmpAvailDisk, errTmpAvail := strconv.ParseFloat(outputDisk["/"][0][:len(outputDisk["/"][0])-1], 64)
	if errTmpAvail != nil {
		fmt.Println("Error at conversion the \"Available\" value for Disk utilization at database load")
	}

	dateDisk := tmpUsedDisk / tmpAvailDisk

	// (MemTotal-MemFree)/MemTotal for RAM

	outputRAM, errRAM := UtilizareRAM()
	if errRAM != nil {
		fmt.Println("Failed to execute the RAM script for database load")
		return errDisk
	}

	tmpFreeRam, errTmpFreeRam := strconv.ParseFloat(outputRAM["Free"][:len(outputRAM["Free"])-1], 64)
	if errTmpFreeRam != nil {
		fmt.Println("Error at conversion the \"Free\" value for RAM utilization at database load")
	}

	tmpTotalRam, errTmpTotalRam := strconv.ParseFloat(outputRAM["Total"][:len(outputRAM["Total"])-1], 64)
	if errTmpTotalRam != nil {
		fmt.Println("Error at conversion the \"Total\" value for RAM utilization at database load")
	}

	dateRAM := (tmpTotalRam - tmpFreeRam) / tmpTotalRam

	// Load average 1 min/nr_processors

	outputCPU, errCPU := UtilizareCPU()
	if errCPU != nil {
		fmt.Println("Failed to execute the CPU script for database load")
		return errDisk
	}

	tmpLoad1min, errTmpLoad1min := strconv.ParseFloat(outputCPU["1min"], 64)
	if errTmpLoad1min != nil {
		fmt.Println("Error at conversion the \"1min\" value for CPU utilization at database load")
	}

	tmpNrProc, errTmpNrProc := strconv.ParseFloat(outputCPU["noProcessors"], 64)
	if errTmpNrProc != nil {
		fmt.Println("Error at conversion the \"noProcessors\" value for CPU utilization at database load")
	}

	dateCPU := new(big.Float)
	dateCPU.SetFloat64(tmpLoad1min / tmpNrProc)

	_, errTblResurse := db.Exec(
		`INSERT INTO
		tblResurse (
			modelSistem, 
			numarSerialSistem, 
			utilizareCPU, 
			utilizareDisk,
			utilizareRAM
		)
		VALUES (
			(
				SELECT idModel
				FROM tblModel
				WHERE numeModel = ?
			),
			(
				SELECT idSerial
				FROM tblSerial
				WHERE numarSerial = ?
			),
			?,
			?,
			?
		)`,
		string(produs), string(serial), dateCPU.Text('e', 40), fmt.Sprintf("%.17g", dateDisk), fmt.Sprintf("%.17g", dateRAM))

	if errTblResurse != nil {
		return errTblResurse
	}

	return nil
}

func LoadDatabase(db *sql.DB, rootpass string) error {

	model, errModel := LoadTblModel(db)
	if errModel != nil {
		return errModel
	}

	producator, errProducator := LoadTblProducator(db)
	if errProducator != nil {
		return errProducator
	}

	procesor, errProcesor := LoadTblProcesor(db)
	if errProcesor != nil {
		return errProcesor
	}

	serial, errSerial := LoadTblSerial(db, rootpass)
	if errSerial != nil {
		return errSerial
	}

	serial, errSistem := LoadTblSistem(db, rootpass, procesor, producator, model, serial)
	if errSistem != nil {
		return errSistem
	}

	errRetea := LoadTblPlaciRetea(db, model, serial)
	if errRetea != nil {
		return errSistem
	}

	errResurse := LoadTblResurse(db, rootpass, model, serial)
	if errResurse != nil {
		return errResurse
	}

	return nil
}

func TriggerLoadCrontab(rootpass string, IPenv string) {

	output, err := BashExec("/var/lib/licenta/api-licenta/update_db.sh", rootpass, IPenv)
	if err != nil {
		fmt.Println("Fail to trigger the database cronjob")
		return
	}

	fmt.Println(string(output))
}
