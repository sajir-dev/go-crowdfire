package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controller "github.com/sajir-dev/go-crowdfire/controllers"
)

func mapUrls() {
	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "Pong") })
	router.POST("/users", controller.CreateUser)
	router.PUT("/posts/:id", controller.UpdatePost)
	router.POST("/follow", controller.FollowPeople)
	router.GET("/posts", controller.GetPosts)
}
