package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func ConnectDB(user string, pass string, ip string, port int) (*sql.DB, error) {

	fmt.Println("Connecting to the database...")

	dsn := user + ":" + pass + "@" + "(" + ip + ":" + strconv.Itoa(port) + ")/metriciDB"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("An Error appeared while connecting to the database...")
		return nil, err
	}

	fmt.Println("Connected to the database.")
	return db, nil
}

func initDB(user string, pass string, ip string, port int) *sql.DB {

	db, errdb := ConnectDB(user, pass, ip, port)
	if errdb != nil {
		log.Fatal(errdb)
	}
	// defer db.Close()

	return db

}

func LoadDatabase(db *sql.DB, rootpass string) error {

	// Preluam informatiile de sistem
	produs, errProdus := bashExec("/var/lib/licenta/api-licenta/get_nume_produs_sistem.sh")
	if errProdus != nil {
		return errProdus
	}
	fmt.Println(string(produs))

	serial, errSerial := bashExec("/var/lib/licenta/api-licenta/get_numar_serial_sistem.sh", rootpass)
	if errSerial != nil {
		return errSerial
	}
	fmt.Println(string(serial))

	furnizor, errFurnizor := bashExec("/var/lib/licenta/api-licenta/get_furnizor_sistem.sh")
	if errFurnizor != nil {
		return errFurnizor
	}
	fmt.Println(string(furnizor))

	procesor, errProcesor := bashExec("/var/lib/licenta/api-licenta/get_procesor_sistem.sh")
	if errProcesor != nil {
		return errProcesor
	}
	fmt.Println(string(procesor))

	// Introducem/Updatam datele de sistem in tabelele foloste de tblSistem
	_, errTblProducator := db.Exec(
		`INSERT INTO
		tblProducator (numeProducator)
		VALUES
		(?)`,
		string(furnizor))

	if errTblProducator != nil {
		if errTblProducator.(*mysql.MySQLError).Number == 1602 {
			fmt.Println("Producatorul se afla deja in baza de date")
		} else {
			return errTblProducator
		}
	}

	_, errTblModel := db.Exec(
		`INSERT INTO
		tblModel (numeModel)
		VALUES (?)`,
		string(produs))

	if errTblModel != nil {
		if errTblModel.(*mysql.MySQLError).Number == 1602 {
			fmt.Println("Modelul se afla deja in baza de date")
		} else {
			return errTblModel
		}
	}

	_, errTblProcesor := db.Exec(
		`INSERT INTO
		tblProcesor (numeProcesor)
		VALUES
		(?)`,
		string(procesor))

	if errTblProcesor != nil {
		if errTblProcesor.(*mysql.MySQLError).Number == 1602 {
			fmt.Println("Procesorul se afla deja in baza de date")
		} else {
			return errTblProcesor
		}
	}

	_, errTblSistem := db.Exec(
		`INSERT INTO
		tblSistem (numarSerial,modelProcesor,producatorSistem,modelSistem)
		VALUES (
			?,
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
		if errTblSistem.(*mysql.MySQLError).Number == 1602 {
			fmt.Println("Sistemul se afla deja in baza de date")
		} else {
			return errTblSistem
		}
	}

	placi_retea, errNIC := bashExec("/var/lib/licenta/api-licenta/get_placi_retea.sh")
	if errNIC != nil {
		return errNIC
	}

	NICs := strings.Split(string(placi_retea), "\n")
	NICs = NICs[:len(NICs)-1]

	fmt.Println(len(NICs))

	for i := 0; i < len(NICs); i++ {

		fmt.Println(NICs[i])

		stare_nic, errStareNIC := bashExec("/var/lib/licenta/api-licenta/get_stare_placa_retea.sh", NICs[i])
		if errStareNIC != nil {
			return errStareNIC
		}
		fmt.Println(string(stare_nic))

		tx_nic, errTxNIC := bashExec("/var/lib/licenta/api-licenta/get_date_transmise_placa_retea.sh", NICs[i])
		if errTxNIC != nil {
			return errTxNIC
		}
		fmt.Println(string(tx_nic))

		rx_nic, errRxNIC := bashExec("/var/lib/licenta/api-licenta/get_date_receptionate_placa_retea.sh", NICs[i])
		if errRxNIC != nil {
			return errRxNIC
		}
		fmt.Println(string(rx_nic))

		dropped_nic, errDroppedNIC := bashExec("/var/lib/licenta/api-licenta/get_date_aruncate_placa_retea.sh", NICs[i])
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
				?,
				?,
				?,
				?,
				?,
				?
			)`,
			string(produs), string(serial), string(NICs[i]), string(stare_nic), string(dropped_nic), string(rx_nic), string(tx_nic))

		if errTblPlaciRetea != nil {
			fmt.Println(errTblPlaciRetea)
		}
	}

	return nil
}
