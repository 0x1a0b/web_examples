package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"github.com/rs/cors"
)


func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/test", Testhandler)
	}
	c := cors.AllowAll()

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))

}

func Testhandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "Success Connection OK"})
}
