package main

import (
	"database/sql"
	"fmt"
	"net/mail"
	"os"
	"path/filepath"
)

func listDir(db *sql.DB, dir string) error {
	return filepath.Walk(dir, walk)
}

func walk(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	msg, err := mail.ReadMessage(file)
	if err != nil {
		return err
	}
	fmt.Println(msg.Header["Subject"])
	return nil
}
