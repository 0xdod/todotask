package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var flagAddr string
var flagDSN string

func main() {
	// inmemStore := NewInMemStore()
	setUpInit()
	store, err := NewPgStore(getDSN())
	if err != nil {
		panic(err)
	}
	server := NewServer(store)
	addr := getAddr()
	log.Printf("Server started locally on port %s", addr)
	if err := http.ListenAndServe(addr, server); err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}
}

func setUpInit() {
	flag.CommandLine.StringVar(&flagAddr, "addr", ":8080", "server address")
	flag.CommandLine.StringVar(&flagDSN, "dsn", "", "dsn")
	flag.Parse()
}

func normalizePort(p string) string {
	if strings.HasPrefix(p, ":") {
		return p
	}
	return ":" + p
}

func getAddr() string {
	if flagAddr != ":8080" {
		return normalizePort(flagAddr)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = flagAddr
	}
	return normalizePort(port)
}

func getDSN() string {
	if flagDSN != "" {
		return flagDSN
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")
	pgPort := os.Getenv("PGPORT")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos",
		host, user, password, dbname, pgPort)
}
