package controller

import (
	"app/matchingAppMatchingService/common/dataStructures"
	"app/matchingAppMatchingService/common/dbInterface"
	"errors"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetAllMatches(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		matches, err := dbInterface.GetAllMatches(db)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		context.IndentedJSON(http.StatusOK, matches)
	}

	return gin.HandlerFunc(handler)
}

func GetMatchById(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		matchId := context.Param("matchid")
		users, err := dbInterface.GetMatchById(db, matchId)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		context.IndentedJSON(http.StatusOK, users)
	}
	return gin.HandlerFunc(handler)
}

func GetAllMatchesForUser(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		id := context.Param("id")
		idInt, errConv := strconv.Atoi(id)
		if errConv != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Server Error!",
			})
			return
		}

		matches, err := dbInterface.GetAllMatchesForUser(db, idInt)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Server Error!",
			})
			return
		}
		if matches == nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Matches not found!",
			})
			return
		}
		context.IndentedJSON(http.StatusOK, matches)
	}

	return gin.HandlerFunc(handler)
}

func DeleteMatch(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var toFind struct {
			matchId string
		}
		errExtract := context.Bind(&toFind)
		if errExtract != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		matchToDelete, errFind := dbInterface.GetMatchById(db, toFind.matchId)
		if errFind != nil {
			context.AbortWithError(http.StatusNotFound, errFind)
			return
		}

		if errDelete := dbInterface.DeleteMatch(db, matchToDelete); errDelete != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": errDelete,
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "Match deleted!",
		})
	}
	return gin.HandlerFunc(handler)
}

func CreateMatch(redis *redis.Client, db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newMatch dataStructures.Match
		if err := context.BindJSON(&newMatch); err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		// Check for confirmed match
		user2LikedUser1, errUser2 := dbInterface.HasUserLiked(redis, &newMatch.LikedId, &newMatch.LikerId)

		if errUser2 != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": errUser2,
			})
			return
		}

		if !user2LikedUser1 {
			context.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": "User didn't like each other!",
			})
			return
		}

		// Create Match in DB
		match, errCreate := dbInterface.CreateMatch(db, &newMatch)
		if errCreate != nil {
			context.AbortWithError(http.StatusInternalServerError, errCreate)
			return
		}
		context.IndentedJSON(http.StatusOK, match)
	}
	return gin.HandlerFunc(handler)
}

func ProposeUser(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		//TODO: Überprüfung auf Like/Dislike

		var possibleUsers []dataStructures.User
		if err := context.BindJSON(&possibleUsers); err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var search dataStructures.Search

		possibleUsers, errPossible := dbInterface.PossibleUsers(db, search.Skill)
		if errPossible != nil {
			context.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": "No fitting users with searched skill found!",
			})
			return
		}

		userToPropose, errProposal := dbInterface.FilterPeople(db, &search, possibleUsers)
		if errProposal != nil {
			context.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": "No fitting users for your search found!",
			})
			return
		}

		context.IndentedJSON(http.StatusOK, userToPropose)
	}
	return gin.HandlerFunc(handler)
}

// Helper

func IsUserOnline() {

}

func CreateMatchAfterLike(redis *redis.Client, matchData *dataStructures.Like) (*dataStructures.Match, error) {

	// Check for confirmed match
	user2LikedUser1, errUser2 := dbInterface.HasUserLiked(redis, &matchData.LikedId, &matchData.LikerId)

	if errUser2 != nil {
		return nil, errUser2
	}

	if !user2LikedUser1 {
		return nil, errors.New("No match yet")
	}

	newMatch := *&dataStructures.Match{
		LikerId: matchData.LikerId,
		LikedId: matchData.LikedId,
	}

	// Create Match in DB
	match, errCreate := dbInterface.CreateMatch(DB, &newMatch)
	if errCreate != nil {
		return nil, errors.New(errCreate.Error())
	}
	return match, nil
}

func FilterPeople(db *gorm.DB, search *dataStructures.Search, possibleUsers []dataStructures.User) ([]dataStructures.User, error) {

	filteredUsers, err := dbInterface.FilterPeople(db, search, possibleUsers)
	if err != nil {
		return nil, err
	}

	return filteredUsers, nil
}
