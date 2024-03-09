package main

func main() {

	// De citit dintr-un fisiere de configurare .ini
	userdb := "admin"
	passdb := "licenta"
	portdb := 3306
	rootpass := "licenta"

	ip := "192.168.0.158"
	port := 8080

	db := initDB(userdb, passdb, portdb)
	initRouters(ip, port, rootpass, db)
}
