package controller

import (
	"app/matchingAppMatchingService/common/dataStructures"
	"app/matchingAppMatchingService/common/dbInterface"
	"app/matchingAppMatchingService/connector"

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
		id := context.Param("id")

		errFind := dbInterface.DeleteMatch(db, id)
		if errFind != nil {
			context.AbortWithError(http.StatusNotFound, errFind)
			return
		}

		/*if errDelete := dbInterface.DeleteMatch(db, matchToDelete); errDelete != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": errDelete,
			})
			return
		}*/

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

		if match != nil {
			_, errLike1 := dbInterface.DeleteLikeEntry(redis, newMatch.LikedId, newMatch.LikerId)
			if errLike1 != nil {
				context.AbortWithError(http.StatusInternalServerError, errLike1)
				return
			}
			_, errLike2 := dbInterface.DeleteLikeEntry(redis, newMatch.LikerId, newMatch.LikedId)
			if errLike2 != nil {
				context.AbortWithError(http.StatusInternalServerError, errLike2)
				return
			}
		}

		context.IndentedJSON(http.StatusAccepted, gin.H{
			"Match created - Id:": match.Id,
		})
	}
	return gin.HandlerFunc(handler)
}

func ProposeUser(db *gorm.DB, redis *redis.Client) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		// Get current User Id
		userId := context.Param("userid")
		convUserId, err := strconv.Atoi(userId)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		searchId := context.Param("id")
		convId, err := strconv.Atoi(searchId)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		search, searchErr := dbInterface.GetSearchById(db, convId)
		if searchErr != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		/*userToFind, errFind := connector.GetProfileById((convId))
		if errFind != nil {
			context.AbortWithError(http.StatusNotFound, errFind)
			return
		}*/

		// Select Users
		selectedSkilledUsers, errUsers := connector.GetProfilesBySkill(search.Skill)
		if errUsers != nil {
			context.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": "No users found!",
			})
			return
		}

		possibleUserToPropose, errProposal := dbInterface.FilterPeople(selectedSkilledUsers, search)
		if errProposal != nil {
			context.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": "No fitting users for your searched settings found!",
			})
			return
		}

		// Check for Dislike
		var userToPropose []dataStructures.User

		for i, _ := range possibleUserToPropose {
			user1DislikedUser2, _ := dbInterface.HasUserDisliked(redis, &convUserId, &possibleUserToPropose[i].ID)

			user2DislikedUser1, _ := dbInterface.HasUserDisliked(redis, &possibleUserToPropose[i].ID, &convUserId)

			if !user1DislikedUser2 {
				if !user2DislikedUser1 {

					user1LikedUser2, _ := dbInterface.HasUserLiked(redis, &convUserId, &possibleUserToPropose[i].ID)

					if !user1LikedUser2 {
						userToPropose = append(userToPropose, possibleUserToPropose[i])
					}
				}
			}
		}

		context.IndentedJSON(http.StatusOK, userToPropose)
	}
	return gin.HandlerFunc(handler)
}

// Helper

func IsUserOnline() {

}

func CreateMatchAfterLike(db *gorm.DB, redis *redis.Client, matchData *dataStructures.Like) (*dataStructures.Match, error) {

	// Check for confirmed match
	user2LikedUser1, errUser2 := dbInterface.HasUserLiked(redis, &matchData.LikedId, &matchData.LikerId)

	if errUser2 != nil {
		return nil, errUser2
	}

	if !user2LikedUser1 {
		return nil, errors.New("No match yet")
	}

	//Check for already existing match
	user1user2Match, errMatch1 := dbInterface.MatchExists(db, &matchData.LikedId, &matchData.LikerId)
	if errMatch1 != nil {
		return nil, errMatch1
	}
	if user1user2Match {
		return nil, errors.New("Match already exists")
	}

	user2user1Match, errMatch2 := dbInterface.MatchExists(db, &matchData.LikerId, &matchData.LikedId)
	if errMatch2 != nil {
		return nil, errMatch1
	}
	if user2user1Match {
		return nil, errors.New("Match already exists")
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

	_, errLike1 := dbInterface.DeleteLikeEntry(redis, matchData.LikedId, matchData.LikerId)
	if errLike1 != nil {
		return nil, errors.New(errCreate.Error())
	}
	_, errLike2 := dbInterface.DeleteLikeEntry(redis, matchData.LikerId, matchData.LikedId)
	if errLike2 != nil {
		return nil, errors.New(errCreate.Error())
	}

	return match, nil
}

func FilterPeople(db *gorm.DB, search *dataStructures.Search, possibleUsers *[]dataStructures.User) ([]dataStructures.User, error) {

	filteredUsers, err := dbInterface.FilterPeople(possibleUsers, search)
	if err != nil {
		return nil, err
	}

	return filteredUsers, nil
}

func ExploreUser(redis *redis.Client) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		// Get current User Id
		userId := context.Param("id")
		convUserId, err := strconv.Atoi(userId)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		var userToPropose []dataStructures.User

		// Get People

		possibleUserIdsToPropose, errUsers := dbInterface.GetAllLikers(redis, &userId)
		if errUsers != nil || possibleUserIdsToPropose == nil {
			context.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": "No users found!",
			})
			return
		}

		/*possibleUserToPropose, errUsers := connector.GetAllProfiles()
		if errUsers != nil {
			context.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": "No users found!",
			})
			return
		}*/

		// Check for Dislike
		for _, otherUserIdString := range *possibleUserIdsToPropose {

			otherUserId, err := strconv.Atoi(otherUserIdString)
			if err != nil {
				context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
				return
			}

			otherUserData, errData := connector.GetProfileById(otherUserId)
			if errData != nil {
				context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
				return
			}

			user1DislikedUser2, _ := dbInterface.HasUserDisliked(redis, &convUserId, &otherUserId)

			user2DislikedUser1, _ := dbInterface.HasUserDisliked(redis, &otherUserId, &convUserId)

			if !user1DislikedUser2 {
				if !user2DislikedUser1 {

					// Check for Like
					user2LikedUser1, _ := dbInterface.HasUserLiked(redis, &otherUserId, &convUserId)

					if user2LikedUser1 {
						userToPropose = append(userToPropose, *otherUserData)
					}
				}
			}
		}

		context.IndentedJSON(http.StatusOK, userToPropose)
	}
	return gin.HandlerFunc(handler)
}
