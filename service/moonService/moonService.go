package moonService

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
)

func Admire(ctx *gin.Context) (map[string]interface{}, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})

	val, err := rdb.Get(ctx, "moon").Result()
	if err == redis.Nil {
		fmt.Println("")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("moon", val)
	}
	fmt.Println("moon", val)
	moon := map[string]interface{}{}
	err = json.Unmarshal([]byte(val), &moon)
	if err != nil {
		log.Printf("unmarshal fail with err:[%+v]\n", moon)
		return nil, err
	}
	return moon, nil
}

func ChangeMoon(ctx *gin.Context, moon string) (int, error){
	rdb := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})
	err := rdb.Set(ctx, "moon", moon, 0).Err()
	if err != nil {
		panic(err)
	}
	return 1, nil
}