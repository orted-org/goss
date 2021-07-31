package service

import (
	"errors"
	"os"
	"time"

	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

var RedisClient *redis.Client = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_URL"),
	Password: "",
	DB:       0,
})

func CreateSession(session string, ttl time.Duration) (string, error) {
	// creating a uuid for session
	uuid := uuid.NewV4()
	// storing the session into redis store
	err := RedisClient.Set(uuid.String(), session, ttl).Err()
	if err != nil {
		return "", nil
	}
	return uuid.String(), nil
}
func GetSession(sessionId string) (string, error) {

	// checking if session id is present
	if len(sessionId) == 0 {
		return "", errors.New("session id missing from query parameter")
	}
	// getting the data from redis store
	res, err := RedisClient.Get(sessionId).Result()
	if err != nil {
		return "", errors.New("session not found")
	}
	return res, nil
}

func DeleteSession(sessionId string) error {

	// checking if session id is present
	if len(sessionId) == 0 {
		return errors.New("no session id provided")
	}
	// removing session from the redis store
	err := RedisClient.Del(sessionId).Err()
	if err != nil {
		return errors.New("session could not be deleted")
	}
	// session deleted and returning nil error
	return nil
}
