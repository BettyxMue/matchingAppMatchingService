package main

import (
	"database/sql"
	"log"

	"app/matchingAppMatchingService/common/database"
	"app/matchingAppMatchingService/controller"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"gorm.io/gorm"
)

var (
	ListenAddr = "localhost:8080"
	RedisAddr  = "localhost:6379"
)

func main() {
	// MySQL Requests
	dbChannel := make(chan *sql.DB)
	gdbChannel := make(chan *gorm.DB)
	go database.InitalizeConnection(dbChannel, gdbChannel)

	db := <-dbChannel
	gdb := <-gdbChannel

	defer db.Close()

	sessionChannel := make(chan *gocql.Session)
	go database.InitDB(sessionChannel)
	session := <-sessionChannel
	defer session.Close()
	router := gin.Default()

	// Get Requests
	router.GET("/match", controller.GetAllMatches(gdb))
	router.GET("/match/:id", controller.GetMatchById(gdb))
	router.GET("/search", controller.GetAllSearches(gdb))
	router.GET("/search/:id", controller.GetSearchByID(gdb))

	// Put Requests
	router.PUT("/match", controller.CreateMatch(gdb))
	router.PUT("/search", controller.CreateSearch(gdb))

	// Update Requests
	router.PUT("/search/:id", controller.UpdateSearch(gdb))
	router.PUT("/match/:id", controller.UpdateMatch(gdb))

	// Delete Requests
	router.DELETE("/search", controller.DeleteSearch(gdb))
	router.DELETE("/match", controller.DeleteMatch(gdb))

	router.Run("0.0.0.0:8080")

	// Redis Requests
	database, err := database.InitRedis(RedisAddr)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error())
	}

	router2 := initRouter(database)
	router2.Run(ListenAddr)
}

func initRouter(database *database.Database) *gin.Engine {
	r := gin.Default()
	return r
}
