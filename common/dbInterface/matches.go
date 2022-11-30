package dbInterface

import (
	"app/matchingAppMatchingService/common/dataStructures"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllMatches(db *gorm.DB) (*[]dataStructures.Match, error) {
	var matches []dataStructures.Match

	err := db.Model(&dataStructures.Match{}).Preload(clause.Associations).Find(&matches).Error

	if err != nil {
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

func IsUserOnline() {

}

func FilterPeople(db *gorm.DB, search *dataStructures.Search, users []dataStructures.User) ([]dataStructures.User, error) {

	//TODO: People in Like / Dislike Tabelle überprüfen

	errLevel := db.Model(&users).Preload(clause.Associations).Where("level=?", search.Level).Find(&users).Error
	if errLevel != nil {
		return nil, errLevel
	}

	errGender := db.Model(&users).Preload(clause.Associations).Where("gender=?", search.Gender).Find(&users).Error
	if errGender != nil {
		return nil, errGender
	}

	//TODO: Radius
	/*errRadius := db.Model(&users).Preload(clause.Associations).Where("radius=<?", search.Radius).Find(&users).Error
	if errRadius != nil {
		return nil, errRadius
	}*/

	/*var userIds []int

	for i, s := range users {
		userIds[i] = int(users[i].ID)
	}*/

	return users, nil
}

func PossibleUsers(db *gorm.DB, skillId int) ([]dataStructures.User, error) {
	var users []dataStructures.User

	err := db.Model(&dataStructures.User{}).Preload(clause.Associations).Where("skillId=?", skillId).Find(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}
