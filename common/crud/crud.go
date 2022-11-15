package crud

import (
	"database/sql"
	"log"

	"app/matchingAppMatchingService/common/dataStructures"

	_ "github.com/go-sql-driver/mysql"
)

func AddMatch(match *dataStructures.Match, db *sql.DB) error {

	statement, err := db.Prepare("INSERT INTO `matches`(`id`,`user1`,`user2`)VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, errInsert := statement.Exec(&match.Id, &match.User1, &match.User2)

	if errInsert != nil {
		log.Fatal(errInsert)
		return errInsert
	}

	return nil

}
