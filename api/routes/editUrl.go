package routes

import (
	"net/http"
	"time"

	"github.com/cdeor/url-shortener/api/database"
	"github.com/cdeor/url-shortener/api/models"
	"github.com/gin-gonic/gin"
)

func EditURL(c *gin.Context) {
	shortID := c.Param("shortID")
	var body models.Request

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot parse JSON",
		})
	}

	r := database.CreateClient(0)
	defer r.Close()

	// check if shortID exists in db or not.
	val, err := r.Get(database.Ctx, shortID).Result()

	if err != nil || val == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ShortID does not exists",
		})
	}

	// update the content of the URL and expiry  with shortID

	err = r.Set(database.Ctx, shortID, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to update yje shortened content",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Content updated",
	})

}
