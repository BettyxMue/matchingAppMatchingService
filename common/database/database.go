package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"app/matchingAppMatchingService/common/dataStructures"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitalizeConnection(dbChannel chan *sql.DB, gdbChannel chan *gorm.DB) *sql.DB {
	dsn := "root:root@tcp(0.0.0.0:3306)/golang_docker?parseTime=true"
	gDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(errors.New("Error connecting to mysql"))
	}
	fmt.Println("Database connected!")

	db, errGetDb := gDb.DB()

	if errGetDb != nil {
		fmt.Println(err)
		panic(errors.New("Error getting DB from gorm"))
	}

	errPing := db.Ping()
	if errPing != nil {
		fmt.Println(errPing)
	}
	setupDatabase(gDb)
	//addMockData(gDb)
	dbChannel <- db
	gdbChannel <- gDb
	return db
}

func setupDatabase(db *gorm.DB) {
	db.AutoMigrate(&dataStructures.Search{})
}

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.Background()
)

func InitRedis(address string) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	return &Database{
		Client: client,
	}, nil
}
