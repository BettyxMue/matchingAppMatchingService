package crud

import (
	"database/sql"
	"log"

	"app/matchingAppMatchingService/common/dataStructures"

	_ "github.com/go-sql-driver/mysql"
)

func AddMatch(match *dataStructures.Match, db *sql.DB) error {

	statement, err := db.Prepare("INSERT INTO `matches`(`id`,`userid1`,`userid2`)VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, errInsert := statement.Exec(&match.Id, &match.UserId1, &match.UserId2)

	if errInsert != nil {
		log.Fatal(errInsert)
		return errInsert
	}

	return nil

}
