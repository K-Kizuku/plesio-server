package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	RedisAddress  string
	RedisPassword string
	RedisDB       int
	RedisPoolSize int
)

// .envを呼び出します。
func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	RedisAddress = os.Getenv("REDIS_ADDRESS")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
	RedisDB, err = strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(err)
	}
	RedisPoolSize, err = strconv.Atoi(os.Getenv("REDIS_POOL_SIZE"))
	if err != nil {
		panic(err)
	}
}
