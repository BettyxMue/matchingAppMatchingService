package dbInterface

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func CreateLike(redis *redis.Client, userId *int, liked *int) (bool, error) {
	//res,err := redis.Do("SADD", userId, liked)
	ress := redis.SAdd(strconv.Itoa(*userId), *liked)
	if ress.Err() != nil {
		return false, ress.Err()
	}
	return true, nil
}

func HasUserLiked(redis *redis.Client, userId1 *int, userId2 *int) (bool, error) {
	result := redis.SIsMember(strconv.Itoa(*userId1), *userId2)

	if result.Err() != nil {
		return false, result.Err()
	}

	return result.Val(), nil
}

func Dislike(redis *redis.Client, userId1 *int, userId2 *int) (bool, error) {
	result := redis.SAdd("dislike"+strconv.Itoa(*userId1), *userId2)
	if result.Err() != nil {
		return false, result.Err()
	}

	var temp = redis.TTL("dislike" + strconv.Itoa(*userId1)).Val()

	if temp == -1000000000 {
		setExpiry := redis.Expire("dislike"+strconv.Itoa(*userId1), time.Duration(7*24*time.Hour))

		if setExpiry.Err() != nil {
			return false, setExpiry.Err()
		}
	}

	return true, nil
}

func HasUserDisliked(redis *redis.Client, userId1 *int, userId2 *int) (bool, error) {
	result := redis.SIsMember("dislike"+strconv.Itoa(*userId1), *userId2)

	if result.Err() != nil {
		return false, result.Err()
	}

	return result.Val(), nil
}

/*
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
}*/

func DeleteLikeEntry(redis *redis.Client, userId1 int, userId2 int) (bool, error) {
	result := redis.Del(strconv.Itoa(userId1), strconv.Itoa(userId2))

	if result.Err() != nil {
		return false, result.Err()
	}

	return true, nil
}
