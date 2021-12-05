package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	redisHost := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":6379",
		Password: redisPassword, // no password set
		DB:       0,             // use default DB
	})

	r := gin.Default()

	r.StaticFile("/", "./index.html")

	r.Static("/static", "./static")

	r.GET("/votes", func(c *gin.Context) {
		catVotes := getVotes(ctx, rdb, "votes.cats")
		dogVotes := getVotes(ctx, rdb, "votes.dogs")
		c.JSON(http.StatusOK, gin.H{
			"cats": catVotes,
			"dogs": dogVotes,
		})
	})

	r.POST("/vote/:animal", func(c *gin.Context) {
		animal := c.Param("animal")
		key := "votes." + animal
		incrementVotes(ctx, rdb, key)
	})

	r.POST("/reset_votes", func(c *gin.Context) {
		setVotes(ctx, rdb, "votes.cats", 0)
		setVotes(ctx, rdb, "votes.dogs", 0)
	})

	r.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func getVotes(ctx context.Context, rdb *redis.Client, key string) int {
	val, err := rdb.Get(ctx, key).Result()
	fmt.Printf("Get", key, val)
	if err != nil {
		fmt.Printf(err.Error())
	}
	valInt, err := strconv.Atoi(val)
	return valInt
}

func incrementVotes(ctx context.Context, rdb *redis.Client, key string) {
	res, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("Incremented", key, "to", res)
}

func setVotes(ctx context.Context, rdb *redis.Client, key string, val int) {
	res, err := rdb.Set(ctx, key, val, 0).Result()
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("Set", key, "to", res)
}
