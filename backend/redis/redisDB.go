package redis
import (
	"fmt"
	"os"
	"log"
	"sync"
	// "strconv"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)
type db struct {
}
type dbInterface interface{
	Conn()
	
}


var err error
var RedisDB *redis.Client
var redisInstance *db

var DBSingletonMux = new(sync.Mutex)

func RedisSingleton() *db {
	if redisInstance == nil {
		DBSingletonMux.Lock()
		if redisInstance == nil {
			redisInstance = &db{}
			defer DBSingletonMux.Unlock()
		}
	}
	return redisInstance
}

func(dbObj *db) Conn() { // 實體化redis.Client 並返回實體的位址
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