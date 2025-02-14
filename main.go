package main

import (
	"fmt"
	"io"
	"main/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", string(body))
	r := gin.Default()
	r.GET("/fetch", handlers.HandleFetch)
	r.Run(":8080")

}
