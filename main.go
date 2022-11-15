package main

import (
	"net/http"

	"app/matchingAppMatchingService/common/dataStructures"
	"app/matchingAppMatchingService/common/database"
	"app/matchingAppMatchingService/common/mockData"

	"app/matchingAppMatchingService/query"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func addMatch(context *gin.Context) {
	var newMatch dataStructures.Match
	if err := context.BindJSON(&newMatch); err != nil {
		return
	}
	mockData.MatchData = append(mockData.Matchdata, newMatch)
	context.IndentedJSON(http.StatusCreated, newMatch)
}

func main() {
	go database.InitalizeConnection()

	router := gin.Default()
	router.GET("/profile", query.GetAllMatches)
	router.GET("/profile/:id", query.GetMatchById)
	router.PUT("/profile", addMatch)
	router.Run("0.0.0.0:8080")
}
