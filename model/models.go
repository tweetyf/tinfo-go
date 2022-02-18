/*
Model:

We have 3 types of storage in this model.

1. go-cache : pure memory cache, for temperary storage, eg session, captcha

2. redis: high speed cache storage, can be deleted and rebuilt as needed

3. mongodb: main database storage

*/
package model

import (
	"context"
	"time"
	"tinfo-go/utils"

	"github.com/go-redis/redis/v8"
	"github.com/patrickmn/go-cache"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	SSID_TIMEOUT     = 24 * time.Hour
	CACHE_TIMEOUT    = 30 * time.Minute
	CACHE_CLEANCYCLE = 10 * time.Minute
	// mongodb url info
	MONGODB_HOST   = "localhost"
	MONGODB_PORT   = "27017"
	MONGODB_DBNAME = "tis"
	// DB collections name
	DBCOLN_USERS     = "users"
	DBCOLN_CHANS     = "channels"
	DBCOLN_POSTS     = "posts"
	DBCOLN_INDEPAGES = "indepages"
	DBCOLN_PREFS     = "prefs"
	// redis info
	REDIS_CACHE_HOST     = "localhost"
	REDIS_CACHE_PORT     = "6379"
	REDIS_CACHE_DB       = 1
	REDIS_CACHE_PASSWORD = ""
)

// Memory-only cache
var CAPTCHA_CACHE = cache.New(CACHE_TIMEOUT, CACHE_CLEANCYCLE)
var SSID_CACHE = cache.New(SSID_TIMEOUT, CACHE_CLEANCYCLE)

// DB instances
var DB *mongo.Database
var REDISDB *redis.Client
var CTX = context.Background()

func InitDatabase() {
	utils.LogD("Initiating Databases...")
	//1. Connecting Mongodb
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + MONGODB_HOST + ":" + MONGODB_PORT))
	if err != nil {
		utils.LogE("%v", err)
	}
	ctx, _ := context.WithTimeout(CTX, 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		utils.LogE("%v", err)
	}
	DB = client.Database(MONGODB_DBNAME)
	//2. Connecting Redis Server
	REDISDB = redis.NewClient(&redis.Options{
		Addr:     REDIS_CACHE_HOST + ":" + REDIS_CACHE_PORT,
		Password: REDIS_CACHE_PASSWORD,
		DB:       REDIS_CACHE_DB,
	})
}

// Redis storage ==============================================
func GetCacheRedis(name string, defaultValue interface{}) interface{} {
	val1, _ := REDISDB.Exists(CTX, name).Result()
	if val1 != 0 {
		val, _ := REDISDB.Get(CTX, name).Result()
		return val
	} else {
		return defaultValue
	}
}

// expire: 0 means never expire
func SetCacheRedis(name string, value interface{}, expire time.Duration) {
	REDISDB.Set(CTX, name, value, expire)
}

// delete from cache
func DelCacheRedis(name string) {
	REDISDB.Del(CTX, name)
}

// Warning: dont use this unless U know what you are doing
func FlushCacheRedis() {
	REDISDB.FlushDB(CTX)
}
