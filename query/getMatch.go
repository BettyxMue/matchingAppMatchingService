package query

import (
	"app/matchingAppMatchingService/common/dataStructures"
	"app/matchingAppMatchingService/common/mockData"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func queryMatches(id string) (*dataStructures.Match, error) {
	for counter, value := range mockData.MatcheData {
		if value.Id == id {
			return &mockData.MatchData[counter], nil
		}
	}
	return &dataStructures.Match{}, errors.New("Match not found!")
}

func GetAllMatches(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, mockData.MatchData)
}

func GetMatchById(context *gin.Context) {
	id := context.Param("id")
	searchedMatch, error := queryMatches(id)
	if error != nil {
		context.AbortWithStatus(http.StatusNotFound)
	}
	context.IndentedJSON(http.StatusOK, searchedMatch)
}
