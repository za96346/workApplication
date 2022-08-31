package redis
import (
	"fmt"
	"os"
	"log"
	// "strconv"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

var err error
var RedisDB *redis.Client

func RedisDBConn() { // 實體化redis.Client 並返回實體的位址
	err = godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	redisIp := os.Getenv("REDIS_DB_IP")
	redisPort := os.Getenv("REDIS_DB_PORT")
	redisPassword := os.Getenv("REDIS_DB_PASSWORD")
	// redisPool := os.Getenv("REDIS_DB_POOL")
	

	RedisDB = redis.NewClient(&redis.Options{
		Addr: redisIp + ":" + redisPort,
		Password: redisPassword, // no password set
		DB: 0,  // use default DB
	})

	pong, err := RedisDB.Ping().Result()
	fmt.Println(pong, err)
}