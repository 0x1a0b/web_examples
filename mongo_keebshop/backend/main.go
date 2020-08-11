package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	return t.Format(time.RFC822)
}

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", ControllerPing)
	}

	c := cors.AllowAll()

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))

}

func ControllerPing(c *gin.Context) {
	now := time.Now()
	c.JSON(http.StatusOK, gin.H{"status": "ok", "time": now.String()})
}
