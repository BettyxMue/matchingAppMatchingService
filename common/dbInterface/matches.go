package dbInterface

import (
	"app/matchingAppMatchingService/common/dataStructures"

	"github.com/gocql/gocql"
)

func GetAllMatches(session *gocql.Session) (*[]dataStructures.Match, error) {
	var match dataStructures.Match
	var matches []dataStructures.Match

	cnqlQuery := "SELECT * FROM match_space.match"
	iterator := session.Query(cnqlQuery).Iter()
	for iterator.Scan(&match.UserId1, &match.UserId2, &match.CreatedAt, &match.UpdatedAt) {
		matches = append(matches, match)
	}

	if errIterator := iterator.Close(); errIterator != nil {
		return nil, errIterator
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
