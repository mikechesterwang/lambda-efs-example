package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/mattn/go-sqlite3"
)

func entry() (string, error) {
	dbpath := os.Getenv("DATA_PATH")
	if dbpath == "" {
		dbpath = "./local.db"
	}
	conn, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return "", err
	}

	defer conn.Close()
	_, err = conn.Exec("CREATE TABLE IF NOT EXISTS v1(c int)")
	if err != nil {
		return "", err
	}

	_, err = conn.Exec("INSERT INTO v1 VALUES (3);")
	if err != nil {
		return "", err
	}

	_, err = conn.Exec("INSERT INTO v1 VALUES (2);")
	if err != nil {
		return "", err
	}

	rows, err := conn.Query("SELECT COUNT(*) FROM v1")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	rtn := ""
	for rows.Next() {
		var c int
		rows.Scan(&c)
		rtn += fmt.Sprint(c)
	}
	return rtn, nil
}

func main() {
	mode := os.Getenv("MODE")
	fmt.Println(mode)
	if mode == "test" {
		r, err := entry()
		if err != nil {
			panic(err)
		}
		fmt.Println(r)
	} else {
		lambda.Start(entry)
	}
}
