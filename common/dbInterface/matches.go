package dbInterface

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"app/matchingAppMatchingService/common/dataStructures"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllMatches(db *gorm.DB) (*[]dataStructures.Match, error) {
	var matches []dataStructures.Match

	err := db.Model(&dataStructures.Match{}).Preload(clause.Associations).Find(&matches).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &matches, nil
}

func GetAllMatchesForUser(db *sql.DB, userId string) (*[]dataStructures.Match, error) {
	rows, err := db.Query("SELECT * FROM match_space.matches WHERE userid1=" + userId + "'")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var matches []dataStructures.Match
	for rows.Next() {
		var match dataStructures.Match
		if errLine := rows.Scan(&match.Id, &match.UserId1, &match.UserId2, &match.ConfirmUser1, &match.ConfirmUser2, &match.SearchId); errLine != nil {
			fmt.Println(errLine)
			return nil, errLine
		}
		matches = append(matches, match)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &matches, nil
}

func GetMatchById(db *gorm.DB, matchId string) (*dataStructures.Match, error) {
	var matches dataStructures.Match

	err := db.Model(&dataStructures.Match{}).Preload(clause.Associations).Where("matchid=?", matchId).Find(&matches).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &matches, nil
}

func UpdateMatch(db *gorm.DB, matchId string, newData *dataStructures.Match) (*dataStructures.Match, error) {
	matchToUpdate, errFind := GetMatchById(db, matchId)
	if errFind != nil {
		return nil, errFind
	}

	changedUser := updateValuesForMatch(matchToUpdate, newData, db)

	result := db.Save(&changedUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return changedUser, nil
}

func DeleteMatch(db *gorm.DB, match *dataStructures.Match) error {
	result := db.Delete(&match)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateMatch(db *sql.DB, userId1 int, userId2 int) error {
	cqlQuery := "INSERT INTO match_space.matches (matchid, userid1, userid2, confirm_user1, confirm_user2, changedat, createdat) VALUES (?,?,?,?,?,?,?) IF NOT EXISTS"
	matchUUID, errUUID := GenerateUUID()
	timeNow := time.Now()
	if errUUID != nil {
		log.Println("Could not generate id for match!")
		return errUUID
	}
	err := db.QueryRow(matchUUID, cqlQuery, userId1, userId2, true, true, timeNow, timeNow).Err()
	if err != nil {
		return err
	}
	return nil
}

func ProposeMatchAgain() {

}

func updateValuesForMatch(oldMatch *dataStructures.Match, newMatch *dataStructures.Match, db *gorm.DB) *dataStructures.Match {
	oldMatch.ConfirmUser1 = newMatch.ConfirmUser1
	oldMatch.ConfirmUser2 = newMatch.ConfirmUser2
	return oldMatch
}

func IsUserOnline() {

}

func GenerateUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
