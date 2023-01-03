package controller

import (
	"app/matchingAppMatchingService/common/dataStructures"
	"app/matchingAppMatchingService/common/dbInterface"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func CreateLike(redis *redis.Client) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var like *dataStructures.Like
		errBind := context.BindJSON(&like)
		if errBind != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Could not find required fields!",
			})
			return
		}
		created, err := dbInterface.CreateLike(redis, &like.LikerId, &like.LikedId)
		if !created {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		match, matchErr := CreateMatchAfterLike(redis, like)

		if matchErr != nil {
			context.JSON(http.StatusCreated, match)
			return
		}

		context.JSON(http.StatusCreated, gin.H{
			"message": "Like created",
		})
	}
	return gin.HandlerFunc(handler)

}

func HasLiked(redis *redis.Client) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var like *dataStructures.Like
		errBind := context.BindJSON(&like)
		if errBind != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Could not find required fields!",
			})
			return
		}
		liked, err := dbInterface.HasUserLiked(redis, &like.LikerId, &like.LikedId)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"liked": liked,
		})

	}
	return gin.HandlerFunc(handler)
}

func Dislike(redis *redis.Client) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var dislike *dataStructures.Dislike
		errBind := context.BindJSON(&dislike)
		if errBind != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": errBind.Error(),
			})
			return
		}
		disliked, err := dbInterface.Dislike(redis, &dislike.DislikerId, &dislike.DislikedId)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"disliked": disliked,
		})
	}
	return gin.HandlerFunc(handler)
}

func HasDisliked(redis *redis.Client) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var dislike *dataStructures.Dislike
		errBind := context.BindJSON(&dislike)
		if errBind != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Could not find required fields!",
			})
			return
		}
		disliked, err := dbInterface.HasUserDisliked(redis, &dislike.DislikerId, &dislike.DislikedId)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"disliked": disliked,
		})

	}
	return gin.HandlerFunc(handler)
}
