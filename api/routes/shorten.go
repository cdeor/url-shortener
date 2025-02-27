package routes

import (
	"net/http"
	"os"
	"time"

	"github.com/actgardner/gogen-avro/v10/test/benchmark/models"
	"github.com/cdeor/url-shortener/api/database"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func ShortenURL(c *gin.Context) {

	var body models.Request

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse JSON"})
		return
	}

	r2 := database.CreateClient(1)
	defer r2.Close()

	val, err := r2.Get(database.Ctx, c.ClientIP()).Result()

	if err == redis.Nil {
		_ = r2.Set(database.Ctx, c.ClientIP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	}

}
