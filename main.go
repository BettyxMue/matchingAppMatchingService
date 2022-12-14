package main

import (
	"database/sql"
	"log"

	"app/matchingAppMatchingService/common/database"
	"app/matchingAppMatchingService/common/initializer"
	"app/matchingAppMatchingService/controller"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	ListenAddr = "localhost:8084"
	RedisAddr  = "localhost:6379"
)

func main() {
	// MySQL Requests
	readyChannel := make(chan bool)
	dbChannel := make(chan *sql.DB)
	gdbChannel := make(chan *gorm.DB)
	redisChannel := make(chan *redis.Client)
	go initializer.LoadEnvVariables(readyChannel)
	<-readyChannel
	log.Println("Service has loaded env")
	go database.InitalizeConnection(dbChannel, gdbChannel)
	go database.InitRedis(RedisAddr, redisChannel)

	db := <-dbChannel
	gdb := <-gdbChannel
	redis := <-redisChannel

	controller.DB = gdb

	defer db.Close()
	defer redis.Close()
	router := gin.Default()

	// Get Requests
	router.GET("/match", controller.GetAllMatches(gdb))
	router.GET("/match/:id", controller.GetMatchById(gdb))
	router.GET("/match/user/:id", controller.GetAllMatchesForUser(gdb))
	router.GET("/search", controller.GetAllSearches(gdb))
	router.GET("/search/:id", controller.GetSearchById(gdb))
	router.GET("/hasLiked", controller.HasLiked(redis)) //Ids mitgeben?
	router.GET("/hasDisliked", controller.HasDisliked(redis))
	router.GET("/searching/:id/:userid", controller.ProposeUser(gdb, redis))
	router.GET("/search/user/:id", controller.GetSearchByUser(gdb))
	router.GET("/exploring/:id", controller.ExploreUser(redis))

	// Put Requests
	router.PUT("/match", controller.CreateMatch(redis, gdb)) // => Können die Aufrufe verkettet werden? BindJSON 2x
	router.PUT("/search", controller.CreateSearch(gdb))
	router.PUT("/like", controller.CreateLike(gdb, redis))
	router.PUT("/dislike", controller.Dislike(redis))

	// Update Requests
	router.PUT("/search/:id", controller.UpdateSearch(gdb))

	// Delete Requests
	router.DELETE("/search/:id", controller.DeleteSearch(gdb))
	router.DELETE("/match/:id", controller.DeleteMatch(gdb))

	router.Run("0.0.0.0:8084")

	// Get Requests
	//router2.GET("/profile/:id/status", controller.IsUserOnline())

	// Put Requests
	//router.PUT("/like", controller.CreateLikeEntry())

	// Update Requests
	//router2.PUT("/like/:id", controller.UpdateLikeEntry())

	// Delete Requests
	//router2.DELETE("/like/:id", controller.DeleteLikeEntry())
}
