package main

func main() {

	// De citit dintr-un fisiere de configurare .ini
	// Database env variables
	userdb := "admin"
	passdb := "licenta"
	ipdb := "localhost"
	portdb := 3306
	rootpass := "licenta"

	// Middleware env variables
	ip := "localhost"
	port := 8080

	db := initDB(userdb, passdb, ipdb, portdb)
	initRouters(ip, port, rootpass, db)
}
