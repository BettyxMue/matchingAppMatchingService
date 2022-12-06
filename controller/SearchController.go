package controller

import (
	"app/matchingAppMatchingService/common/dataStructures"
	"app/matchingAppMatchingService/common/dbInterface"
	"strconv"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllSearches(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		chats, err := dbInterface.GetAllSearches(db)
		if err != nil {
			fmt.Println(err)
			context.AbortWithStatusJSON(http.StatusNoContent, gin.H{
				"error": err,
			})
			return
		}
		context.JSON(http.StatusOK, chats)
	}
	return gin.HandlerFunc(handler)
}

func CreateSearch(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newSearch dataStructures.Search
		if err := context.BindJSON(&newSearch); err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		searchToReturn, errCreate := dbInterface.CreateSearch(db, &newSearch)
		if errCreate != nil {
			context.AbortWithError(http.StatusInternalServerError, errCreate)
			return
		}
		context.IndentedJSON(http.StatusOK, searchToReturn)
	}
	return gin.HandlerFunc(handler)
}

func GetSearchByID(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		id := context.Param("id")

		convSearchId, err := strconv.Atoi(id)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		users, err := dbInterface.GetSearchById(db, convSearchId)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		context.IndentedJSON(http.StatusOK, users)
	}
	return gin.HandlerFunc(handler)
}

func DeleteSearch(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		searchId := context.Param("id")

		convSearchId, err := strconv.Atoi(searchId)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		searchToDelete, findErr := dbInterface.GetSearchById(db, convSearchId)
		if findErr != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Search not found!",
			})
			return
		}

		deleteErr := dbInterface.DeleteSearch(db, searchToDelete)
		if deleteErr != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": deleteErr,
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "Search deleted!",
		})

	}
	return gin.HandlerFunc(handler)
}

func UpdateSearch(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newData *dataStructures.Search
		var searchId = context.Param("searchid")

		convSearchId, err := strconv.Atoi(searchId)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		errBind := context.BindJSON(&newData)
		if errBind != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": errBind,
			})
			return
		}

		updatedSearch, errUpdate := dbInterface.UpdateSearch(db, convSearchId, newData)
		if errUpdate != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": errBind,
			})
			return
		}
		context.JSON(http.StatusOK, updatedSearch)

	}
	return gin.HandlerFunc(handler)
}
