package routes

import (
	"encoding/json"
	"net/http"

	"github.com/cdeor/url-shortener/api/database"
	"github.com/gin-gonic/gin"
)

type TagRequest struct {
	ShortID string `json:"shortID"`
	Tag     string `json:"tag"`
}

func AddTAG(c *gin.Context) {
	var tagReq TagRequest
	if err := c.ShouldBind(&tagReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	shortId := tagReq.ShortID
	tag := tagReq.Tag

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortId).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found for shortID",
		})
		return
	}

	var data map[string]interface{}

	if err := json.Unmarshal([]byte(val), data); err != nil {
		data = make(map[string]interface{})
		data["data"] = val
	}

	// check if tags field exists

	var tags []string
	if existingTags, ok := data["tags"].([]interface{}); ok {
		for _, t := range existingTags {
			if strTag, ok := t.(string); ok {
				tags = append(tags, strTag)
			}
		}
	}

	// check for duplicate tags

	for _, existingTag := range tags {
		if existingTag == tag {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "tag already exists",
			})
			return
		}
	}

	// add the new tag  to the tag slice
	tags = append(tags, tag)
	data["tags"] = tags

	// marshal the updated data back to json

	updatedData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to marshal updated data",
		})
		return
	}

	if err = r.Set(database.Ctx, shortId, updatedData, 0).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update the db",
		})
		return
	}

	// response with updated data
	c.JSON(http.StatusOK, data)

}
