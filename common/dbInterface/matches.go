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

func DeleteMatch(db *gorm.DB, matchId string) error {
	var matchToDelete dataStructures.Match

	err := db.Model(&dataStructures.Match{}).Where("id=?", matchId).Find(&matchToDelete).Error
	if err != nil {
		return err
	}

	result := db.Delete(&matchToDelete)
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

func FilterPeople(users *[]dataStructures.User, search *dataStructures.Search) ([]dataStructures.User, error) {

	var possibleUsers []dataStructures.User

	for i, _ := range *users {

		// Check for Gender
		if search.Gender != 4 {

			if (*users)[i].Gender == search.Gender {

				//TODO: Check for Radius

				possibleUsers = append(possibleUsers, (*users)[i])

				/*var skills = (*users)[i].AchievedSkills
				for j, _ := range skills {

					if skills[j].ID == uint(search.Id) {

						// Check for Level
						if skills[j].Level == search.Level {
							possibleUsers = append(possibleUsers, (*users)[i])
						}
					}
				}*/
			}
		} else {
			//TODO: Check for Radius

			possibleUsers = append(possibleUsers, (*users)[i])

			/*var skills = (*users)[i].
			for j, _ := range skills {

				if skills[j].ID == uint(search.Id) {

					// Check for Level
					if skills[j].Level == search.Level {
						possibleUsers = append(possibleUsers, (*users)[i])
					}
				}
			}*/
		}
	}

	return possibleUsers, nil
}

/*func PossibleUsers(users *[]dataStructures.User, skillId int) ([]dataStructures.User, error) {

	var possbileUsers []dataStructures.User

	for i, _ := range *users {

		var skills = (*users)[i].AchievedSkills
		for j, _ := range skills {

			if skills[j].ID == uint(skillId) {
				possbileUsers = append(possbileUsers, (*users)[i])
			}
		}
	}

	err := db.Model(&dataStructures.User{}).Preload(clause.Associations).Where("skillId=?", skillId).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return possbileUsers, nil
}*/
