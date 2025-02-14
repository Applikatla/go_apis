package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transact struct {
	Transaction string `json:"transaction"`
}

// Define the struct to match the JSON structure
type Response struct {
	Proof bool   `json:"proof"`
	Hex   []any  `json:"hex"`
	Root  Buffer `json:"root"`
}

// Define the Buffer struct to handle {"type":"Buffer","data":[...]}
type Buffer struct {
	Type string `json:"type"`
	Data []byte `json:"data"`
}

func HandlePost(c *gin.Context) {
	myUrl := "http://localhost:8001/mint"
	var t transact
	err := c.ShouldBindJSON(&t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	reBody, err := json.Marshal(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to encode JSON"})
		return
	}

	response, err := http.Post(myUrl, "application/json", bytes.NewBuffer(reBody))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send request"})
		return
	}
	var b Response
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &b)
	// Send the fetched JSON data to the frontend
	c.JSON(http.StatusOK, b)

}
