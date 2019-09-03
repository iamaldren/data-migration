package data_migration

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"runtime/debug"
)

var client *redis.Client

type Params struct {
	Key string
	UpdateForeignId bool
	TargetTable string
	SourceTable string
}

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", //when empty means no password is set in Redis
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}

	fmt.Println("Successfully pinged redis with response of " + pong)
}

func MigrateDataForKVPair(params Params) {
	result, err := client.Get(params.Key).Result()
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}

	student := Student{}
	err = json.Unmarshal([]byte(result), student)
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}

	err = Insert(student)
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}
}

func UpdateForeignId(params Params) {
	fmt.Println(params)
}