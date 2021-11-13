package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
)

type Entity struct {
	Id          int         `db:"id"`
	Name        string      `db:"name"`
	Description string      `db:"description"`
	Properties  string         `db:"properties"`
}

type Amounts struct {
	Amount int64 `json:"amount"`
	ImageUrl string `json:"image_url"`
	Description string `json:"description"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "pass"
	dbname     = "db"
	schemaName = "test"
)

func main() {
	// or you can use this
	// db, err := sql.Open("postgres", "postgres://root:pass@localhost/db?sslmode=disable&search_path=test")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable search_path=%s",
		host, port, user, password, dbname, schemaName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := Entity{Id:1}

	err = db.QueryRow("SELECT name, description, properties FROM entity WHERE id = $1",
		e.Id).Scan(&e.Name, &e.Description, &e.Properties)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Select DB Result before update: %+v\n", e)

	test := map[string][]Amounts{}
	err = json.Unmarshal([]byte(e.Properties), &test)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	test["amounts"][0].Amount = 999992
	test["amounts"][1].Amount = 999993
	res, err := json.Marshal(test["amounts"])
	if err != nil {
		fmt.Println("Error: ", err)
	}

	rest, err := db.Exec("update entity set properties = jsonb_set(\"properties\", '{\"amounts\"}', $1, true) where id = 1", fmt.Sprintf( "%s", res))
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("\nRows Updated %v", fmt.Sprint(rest.RowsAffected()))

	err = db.QueryRow("SELECT name, description, properties FROM entity WHERE id = $1",
		e.Id).Scan(&e.Name, &e.Description, &e.Properties)

	fmt.Printf("\nSelect DB Result after update: %+v\n", e)
}