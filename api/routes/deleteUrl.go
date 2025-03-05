package routes

import (
	"net/http"

	"github.com/cdeor/url-shortener/api/database"
	"github.com/gin-gonic/gin"
)

func DeleteURL(c *gin.Context) {

	shortID := c.Param("shortID")

	r := database.CreateClient(0)
	defer r.Close()

	if err := r.Del(database.Ctx, shortID).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to delete shortened URL",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Shortened URL deleted successfully",
	})

}
