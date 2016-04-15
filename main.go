package main

import (
	"flag"
	"log"
)

func main() {
	dir := flag.String("dir", ".", "mail directory")
	dbfile := flag.String("db", "db", "mail database")
	flag.Parse()
	db, err := newDB(*dir, *dbfile)
	if err != nil {
		log.Fatal(err)
	}
	listDir(db, *dir)
}
