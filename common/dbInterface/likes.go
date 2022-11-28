package dbInterface

import (
	"fmt"

	"app/matchingAppMatchingService/common/dataStructures"

	"github.com/go-redis/redis"
)

func CreateLikeTable() {

}

func CreateLikeEntry(c redis.Conn) error {
	var likeEntry dataStructures.Like
	_, err := c.Do("SET", likeEntry.LikedId, likeEntry.LikerId)
	if err != nil {
		return err
	}
	return nil
}

func GetLikeEntryByLiker(c redis.Conn, liker string) error {
	var s, err = redis.String(c.Do("GET", liker))
	if err == redis.ErrNil {
		fmt.Printf("%s does not exist\n", liker)
	} else if err != nil {
		return err
	} else {
		fmt.Printf("%s = %s\n", liker, s)
	}
	return nil
}

func UpdateLikeEntry() {
	GetLikeEntryByLiker()
}

func DeleteLikeEntry() {

}
