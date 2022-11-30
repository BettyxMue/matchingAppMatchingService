package dbInterface

import (
	"fmt"

	"app/matchingAppMatchingService/common/dataStructures"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateSearch(db *gorm.DB, search *dataStructures.Search) (*dataStructures.Search, error) {
	result := db.Create(&search)

	if result.Error != nil {
		return nil, result.Error
	}

	return search, nil
}

func GetAllSearches(db *gorm.DB) (*[]dataStructures.Search, error) {
	var searches []dataStructures.Search

	err := db.Model(&dataStructures.Search{}).Preload(clause.Associations).Find(&searches).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &searches, nil
}

func GetSearchById(db *gorm.DB, searchId string) (*dataStructures.Search, error) {
	var searches dataStructures.Search

	err := db.Model(&dataStructures.Search{}).Preload(clause.Associations).Where("searchid=?", searchId).Find(&searches).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &searches, nil
}

func UpdateSearch(db *gorm.DB, searchId string, newData *dataStructures.Search) (*dataStructures.Search, error) {
	searchToUpdate, errFind := GetSearchById(db, searchId)
	if errFind != nil {
		return nil, errFind
	}

	changedSearch := updateValuesForSearch(searchToUpdate, newData, db)

	result := db.Save(&changedSearch)

	if result.Error != nil {
		return nil, result.Error
	}

	return changedSearch, nil
}

func DeleteSearch(db *gorm.DB, search *dataStructures.Search) error {
	result := db.Delete(&search)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func updateValuesForSearch(oldSearch *dataStructures.Search, newSearch *dataStructures.Search, db *gorm.DB) *dataStructures.Search {
	oldSearch.Name = newSearch.Name
	oldSearch.Level = newSearch.Level
	oldSearch.Gender = newSearch.Gender
	oldSearch.Skill = newSearch.Skill
	oldSearch.Radius = newSearch.Radius
	return oldSearch
}
