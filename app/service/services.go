package service

import (
	"errors"
	"os"
	"time"

	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

func getRedisAddr() string {
	fromEnv := os.Getenv("REDIS_URL")
	if len(fromEnv) == 0 {
		// default with reference to the docker service
		return "localhost:6379"
	}
	return fromEnv
}

var RedisClient *redis.Client = redis.NewClient(&redis.Options{
	Addr:     getRedisAddr(),
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
func TruncateStore() error {
	err := RedisClient.FlushAll().Err()
	if err != nil {
		return errors.New("could not truncate the seesion store")
	}
	return nil
}
