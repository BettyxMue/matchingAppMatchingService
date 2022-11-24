package main

import (
	"database/sql"

	"app/matchingAppMatchingService/common/database"
	"app/matchingAppMatchingService/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbChannel := make(chan *sql.DB)
	gdbChannel := make(chan *gorm.DB)
	go database.InitializeConnection(dbChannel, gdbChannel)

	db := <-dbChannel
	gdb := <-gdbChannel

	defer db.Close()

	router := gin.Default()
	// Get Requests
	router.GET("/search", controller.GetAllSearches(gdb))
	router.GET("/search/:id", controller.GetSearchByID(gdb))
	router.GET("/match", controller.GetAllMatches(gdb))
	router.GET("/match/:id", controller.GetAllMatchesForUser(gdb))
	router.GET("/match/:id/users", controller.GetAllMatchesForUser(gdb))

	// Put Requests
	router.PUT("/signUp", controller.CreateProfile(gdb))
	router.PUT("/skill", controller.CreateSkill(gdb))
	router.PUT("/login", controller.LoginUser(gdb))

	// Update Requests
	router.PUT("/profile/:id", controller.UpdateUser(gdb))

	// Delete Requests
	router.DELETE("/profile", controller.DeleteUser(gdb))
	router.DELETE("/skill/:id", controller.DeleteSkill(gdb))

	router.Run("0.0.0.0:8080")
}
