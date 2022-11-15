package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"app/matchingAppMatchingService/common/crud"
	"app/matchingAppMatchingService/common/mockData"

	_ "github.com/go-sql-driver/mysql"
)

func InitializeConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(database:3306)/golang_docker")
	if err != nil {
		fmt.Println(err)
		panic(errors.New("Error connecting to mysql"))
	}
	defer db.Close()

	fmt.Println("Database connected!")

	errPing := db.Ping()
	if errPing != nil {
		fmt.Println(errPing)
	}
	createMatchTable(db)
	addMockData(db)
	return db
}

func createMatchTable(db *sql.DB) error {
	fmt.Println("Creating table...")
	query := "CREATE TABLE IF NOT EXISTS matches(id int primary key AUTO_INCREMENT, user1 USER_DEFINED_TYPE_SCHEMA, user2 USER_DEFINED_TYPE_SCHEMA, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)"
	fmt.Println("Sending Command!")
	res, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Command sended!")
	rows, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Rows affected when creating table: %d\n", rows)
	return nil
}

func addMockData(db *sql.DB) {
	err := crud.AddMatch(&mockData.MatchData[0], db)
	if err != nil {
		log.Fatal(err)
	}
}
