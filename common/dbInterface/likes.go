package dbInterface

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func CreateLike(redis *redis.Client, userId *int, liked *int) (bool, error) {
	ress := redis.SAdd(strconv.Itoa(*userId), *liked)
	if ress.Err() != nil {
		return false, ress.Err()
	}

	result := redis.SAdd("liked"+strconv.Itoa(*liked), *userId)
	if result.Err() != nil {
		return false, result.Err()
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

func GetAllLikers(redis *redis.Client, userId1 *int) (*[]string, error) {

	result, err := redis.SMembers("liked" + strconv.Itoa(*userId1)).Result()

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func DeleteLikeEntry(redis *redis.Client, userId1 int, userId2 int) (bool, error) {
	result := redis.SRem(strconv.Itoa(userId1), strconv.Itoa(userId2))

	if result.Err() != nil {
		return false, result.Err()
	}

	return true, nil
}
