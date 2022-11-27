package dbInterface

import (
	"fmt"

	"app/matchingAppMatchingService/common/dataStructures"

	"github.com/gocql/gocql"
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

func GetAllMatchesForUser(session *gocql.Session, userId int) (*[]dataStructures.Match, error) {
	var match dataStructures.Match
	var matches []dataStructures.Match

	cnqlQuery1 := "SELECT * FROM match_space.match WHERE userid1=?"
	cnqlQuery2 := "SELECT * FROM match_space.match WHERE userid2=? ALLOW FILTERING"
	iterator1 := session.Query(cnqlQuery1, userId).Iter()
	iterator2 := session.Query(cnqlQuery2, userId).Iter()
	for iterator1.Scan(&match.UserId1, &match.UserId2, &match.UpdatedAt, &match.CreatedAt) {
		matches = append(matches, match)
	}
	if errIterator1 := iterator1.Close(); errIterator1 != nil {
		return nil, errIterator1
	}
	for iterator2.Scan(&match.UserId1, &match.UserId2, &match.UpdatedAt, &match.CreatedAt) {
		matches = append(matches, match)
	}
	if errIterator2 := iterator2.Close(); errIterator2 != nil {
		return nil, errIterator2
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

func CreateMatch(db *gorm.DB, match *dataStructures.Match) (*dataStructures.Match, error) {
	result := db.Create(&match)

	if result.Error != nil {
		return nil, result.Error
	}

	return match, nil
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
