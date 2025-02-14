package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// HandleFetch fetches data from JSONPlaceholder and sends it to the frontend
func HandleFetch(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	defer resp.Body.Close()

	var d Data
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode JSON"})
		return
	}

	// Send the fetched JSON data to the frontend
	c.JSON(http.StatusOK, d)
}
