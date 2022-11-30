package dbInterface

import (
	"fmt"

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
func GetAllMatchesForUser(db *gorm.DB, userId int) (*[]dataStructures.Match, error) {
	var matches []dataStructures.Match

	err := db.Model(&dataStructures.Match{}).Where("userId1=?", userId).Find(&matches).Error

	if err != nil {
		return nil, err
	}
	return &matches, nil
}

func GetMatchById(db *gorm.DB, matchId string) (*dataStructures.Match, error) {
	var matches dataStructures.Match

	err := db.Model(&dataStructures.Match{}).Where("matchid=?", matchId).Find(&matches).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &matches, nil
}

func DeleteMatch(db *gorm.DB, match *dataStructures.Match) error {
	result := db.Delete(&match)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateMatch(db *gorm.DB, match *dataStructures.Match) (*dataStructures.Match, error) {

	errInsert := db.Save(match)

	if errInsert.Error != nil {
		return nil, errInsert.Error
	}

	return match, nil
}

func ProposeMatchAgain() {

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
