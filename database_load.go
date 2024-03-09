package main

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(user string, pass string, port int) (*sql.DB, error) {

	dsn := user + ":" + pass + "@" + "(localhost:" + strconv.Itoa(port) + ")/metriciDB"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initDB(user string, pass string, port int) {

	db, errdb := ConnectDB(user, pass, port)
	if errdb != nil {
		log.Fatal(errdb)
	}
	defer db.Close()

}

// func LoadDatabase(db *sql.DB) error {

// 	password := "licenta" // de facut cu vault!

// 	produs, err := bashExec("/var/lib/licenta/api/get_nume_produs_sistem.sh")

// 	serial, err := bashExec("/var/lib/licenta/api/get_numar_serial_sistem.sh", password)

// 	furnizor, err := bashExec("/var/lib/licenta/api/get_furnizor_sistem.sh")

// 	procesor, err := bashExec("/var/lib/licenta/api/get_procesor_sistem.sh")

// 	placi_retea, err := bashExec("/var/lib/licenta/api/get_placi_retea.sh")

// 	// stare_nic, err := bashExec("/var/lib/licenta/api/get_stare_placa_retea.sh", param) cu array-ul din placi retea

// 	// tx_nic, err := bashExec("/var/lib/licenta/api/get_date_transmise_placa_retea.sh", param) cu array-ul din placi retea

// 	// rx_nic, err := bashExec("/var/lib/licenta/api/get_date_receptionate_placa_retea.sh", param) cu array-ul din placi retea

// 	// dropped_nic, err := bashExec("/var/lib/licenta/api/get_date_aruncate_placa_retea.sh", param) cu array-ul din placi retea

// 	_, errdb := db.Exec("INSERT INTO table_name (column_name) VALUES (?)", "value")
// 	if errdb != nil {
// 		return err
// 	}
// 	return nil
// }
