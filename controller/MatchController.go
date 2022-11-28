package controller

import (
	"app/matchingAppMatchingService/common/dataStructures"
	"app/matchingAppMatchingService/common/dbInterface"
	"fmt"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"gorm.io/gorm"
)

func GetAllMatches(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		users, err := dbInterface.GetAllMatches(db)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		context.IndentedJSON(http.StatusOK, users)
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

func GetAllMatchesForUser(session *gocql.Session) gin.HandlerFunc {
	handler := func(context *gin.Context) {

		userId, errConv := strconv.Atoi(context.Param("id"))
		if errConv != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "User ID must be a number",
			})
			return
		}
		matches, err := dbInterface.GetAllMatchesForUser(session, userId)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		context.JSON(http.StatusOK, matches)
	}
	return gin.HandlerFunc(handler)
}

func UpdateMatch(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newData *dataStructures.Match
		var matchId = context.Param("matchid")
		errBind := context.BindJSON(&newData)
		if errBind != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": errBind,
			})
			return
		}

		updatedMatch, errUpdate := dbInterface.UpdateMatch(db, matchId, newData)
		if errUpdate != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": errBind,
			})
			return
		}
		context.JSON(http.StatusOK, updatedMatch)

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

func CreateMatch(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newMatch dataStructures.Search
		if err := context.BindJSON(&newMatch); err != nil {
			fmt.Println(err)
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		matchToReturn, errCreate := dbInterface.CreateSearch(db, &newMatch)
		if errCreate != nil {
			fmt.Println(errCreate)
			context.AbortWithError(http.StatusInternalServerError, errCreate)
			return
		}
		context.IndentedJSON(http.StatusOK, matchToReturn)
	}
	return gin.HandlerFunc(handler)
}

func ProposeMatchAgain() {

}

func IsUserOnline() {

}
